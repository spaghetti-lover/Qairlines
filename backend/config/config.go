package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	MailSenderName          string        `mapstructure:"MAIL_SENDER_NAME"`
	MailSenderAddress       string        `mapstructure:"MAIL_SENDER_ADDRESS"`
	MailSenderPassword      string        `mapstructure:"MAIL_SENDER_PASSWORD"`
	RedisAddress            string        `mapstructure:"REDIS_ADDRESS"`
	DBDriver                string        `mapstructure:"DB_DRIVER"`
	DBSource                string        `mapstructure:"DB_SOURCE"`
	ServerAddressPort       string        `mapstructure:"SERVER_ADDRESS_PORT"`
	TokenSymmetricKey       string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration     time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	AppEnv                  string        `mapstructure:"APP_EVN"`
	RateLimiterRequestSec   int           `mapstructure:"RATE_LIMITER_REQUEST_SEC"`
	RateLimiterRequestBurst int           `mapstructure:"RATE_LIMITER_REQUEST_BURST"`
	StripeSecretKey         string        `mapstructure:"STRIPE_SECRET_KEY"`
	RedisDB                 string        `mapstructure:"REDIS_DB"`
	RedisUsername           string        `mapstructure:"REDIS_USERNAME"`
	RedisPassword           string        `mapstructure:"REDIS_PASSWORD"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return

}
