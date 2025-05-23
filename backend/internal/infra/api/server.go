package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spaghetti-lover/qairlines/config"
	db "github.com/spaghetti-lover/qairlines/db/sqlc"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/auth"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/news"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/user"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
	"github.com/spaghetti-lover/qairlines/internal/infra/postgresql"
	"github.com/spaghetti-lover/qairlines/pkg/token"
)

type Server struct {
	store      *db.Store
	router     http.Handler
	tokenMaker token.Maker
}

func NewServer(config config.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	// Health
	healthRepo := postgresql.NewHealthRepositoryPostgres(store)
	healthUseCase := usecases.NewHealthUseCase(healthRepo)
	healthHandler := handlers.NewHealthHandler(healthUseCase)

	// User
	userRepo := postgresql.NewUserRepositoryPostgres(store, tokenMaker)
	userGetAllUseCase := user.NewUserGetAllUseCase(userRepo)
	userCreateUseCase := user.NewUserCreateUseCase(userRepo)
	userGetByEmailUseCase := user.NewUserGetByEmailUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userGetAllUseCase, userCreateUseCase, userGetByEmailUseCase)

	// Auth
	loginUseCase := auth.NewLoginUseCase(userRepo, tokenMaker)
	changePasswordUseCase := auth.NewChangePasswordUseCase(userRepo)
	authHandler := handlers.NewAuthHandler(loginUseCase, changePasswordUseCase)

	// News
	newsRepo := postgresql.NewNewsModelRepositoryPostgres(store)
	newsGetAllUseCase := news.NewNewsGetAllUseCase(newsRepo)
	newsHandler := handlers.NewNewsHandler(newsGetAllUseCase)

	// Use gorilla/mux for routing
	router := mux.NewRouter()

	// Health API
	router.HandleFunc("/health", healthHandler.ServeHTTP).Methods("GET")

	// News API
	router.HandleFunc("/api/news", newsHandler.GetAllNews).Methods("GET")

	// User API
	router.HandleFunc("/api/user", userHandler.CreateUser).Methods("POST")

	// Auth API
	router.HandleFunc("/api/auth/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/api/user/{id}/password", authHandler.ChangePassword).Methods("PUT")

	// Wrap router with CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	server := &Server{
		store:      store,
		router:     corsHandler,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
