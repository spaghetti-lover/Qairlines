package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spaghetti-lover/qairlines/config"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu      sync.RWMutex
	clients = make(map[string]*Client)
)

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}
	return ip
}

func getRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	client, exists := clients[ip]
	if !exists {
		config, err := config.LoadConfig(".")
		if err != nil {
			panic("invalid env: " + err.Error())
		}
		requestSec := config.RateLimiterRequestSec
		burst := config.RateLimiterRequestBurst
		limiter := rate.NewLimiter(rate.Limit(requestSec), burst)
		newClient := &Client{limiter: limiter, lastSeen: time.Now()}
		mu.Lock()
		clients[ip] = newClient
		mu.Unlock()
		return limiter
	}
	client.lastSeen = time.Now()
	return client.limiter
}

func CleanUpClients() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

// Use this command to test "ab -n 20 -c 1 localhost:8080/api/news/all"
func RateLimitingMiddleware(recoveryLogger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIP(ctx)
		limiter := getRateLimiter(ip)
		if !limiter.Allow() {
			if shouldLogRateLimit(ip) {
				recoveryLogger.Warn().
					Str("client_ip", ctx.ClientIP()).
					Str("user_agent", ctx.Request.UserAgent()).
					Str("referer", ctx.Request.Referer()).
					Str("protocol", ctx.Request.Proto).
					Str("host", ctx.Request.Host).
					Str("remote_addr", ctx.Request.RemoteAddr).
					Str("request_uri", ctx.Request.RequestURI).
					Interface("headers", ctx.Request.Header).
					Msg("rate limiter exceeded")
			}
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "Too many requests",
				"message": "Too many requests.Please try again",
			})
			return
		}
		ctx.Next()
	}
}

var rateLimitLogCache = sync.Map{}

const rateLimitLogTTL = 10 * time.Second

func shouldLogRateLimit(ip string) bool {
	now := time.Now()
	if val, ok := rateLimitLogCache.Load(ip); ok {
		if t, ok := val.(time.Time); ok && now.Sub(t) < rateLimitLogTTL {
			return false
		}
	}

	rateLimitLogCache.Store(ip, now)
	return true
}
