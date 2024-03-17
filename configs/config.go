package configs

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppAuthor    string
	AppPort      string
	JWTSecretKey []byte
	JWTDuration  time.Duration
	DatabaseDSN  string
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func InitConfig() *Config {
	config := viper.New()
	config.AutomaticEnv()
	config.SetConfigFile("config.json")
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	return &Config{
		AppAuthor:    config.GetString("app.author"),
		AppPort:      config.GetString("app.port"),
		JWTSecretKey: []byte(config.GetString("jwt.secret_key")),
		JWTDuration:  config.GetDuration("jwt.duration"),
		DatabaseDSN:  config.GetString("database.mysql_dsn"),
		ClientID:     config.GetString("oauth.google.client_id"),
		ClientSecret: config.GetString("oauth.google.client_secret"),
		RedirectURL:  config.GetString("oauth.google.redirect_url"),
	}
}
