package config

import (
	"hedwig/cmd/smtp"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int
}

type Config struct {
	AppName string
	IsDebug bool
	Smtp    smtp.SmtpConfig
	Server  ServerConfig
}

func New() (*Config, error) {
	config := &Config{}

	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	config.AppName = "Bum Store"
	config.IsDebug = viper.GetBool("DEBUG")
	config.Server = ServerConfig{Port: viper.GetInt("PORT")}
	config.Smtp = smtp.SmtpConfig{
		Host:     viper.GetString("SMTP_HOST"),
		Port:     viper.GetString("SMTP_PORT"),
		From:     viper.GetString("SMTP_USERNAME"),
		Password: viper.GetString("SMTP_PASSWORD"),
	}
	return config, nil
}

func (c *Config) Print() {
	_, _ = pp.Println(c)
}
