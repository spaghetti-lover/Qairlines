package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	MailFrom                string        `mapstructure:"MAIL_FROM"`
	MailPassword            string        `mapstructure:"MAIL_PASSWORD"`
	MailPort                string        `mapstructure:"MAIL_PORT"`
	MailHost                string        `mapstructure:"MAIL_HOST"`
	MailServer              string        `mapstructure:"MAIL_SERVER"`
	KafkaBrokerURL          string        `mapstructure:"KAFKA_BROKER_URL"`
	KafkaTopic              string        `mapstructure:"KAFKA_TOPIC"`
	KafkaGroupID            string        `mapstructure:"KAFKA_GROUP_ID"`
	DBDriver                string        `mapstructure:"DB_DRIVER"`
	DBSource                string        `mapstructure:"DB_SOURCE"`
	ServerAddressPort       string        `mapstructure:"SERVER_ADDRESS_PORT"`
	TokenSymmetricKey       string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration     time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	AppEnv                  string        `mapstructure:"APP_EVN"`
	RateLimiterRequestSec   int           `mapstructure:"RATE_LIMITER_REQUEST_SEC"`
	RateLimiterRequestBurst int           `mapstructure:"RATE_LIMITER_REQUEST_BURST"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return

}
