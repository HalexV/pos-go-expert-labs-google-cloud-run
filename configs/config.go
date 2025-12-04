package configs

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type conf struct {
	WebServerPort string `mapstructure:"WEB_SERVER_PORT" validate:"required"`
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY" validate:"required"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg conf

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// ENV
	viper.AutomaticEnv()

	// bind ENV VARS explicitamente
	keys := []string{
		"WEB_SERVER_PORT",
		"WEATHER_API_KEY",
	}
	for _, key := range keys {
		viper.BindEnv(key)
	}

	// Try load .env
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(".env não encontrado — usando apenas variáveis do ambiente")
		} else {
			return nil, err
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
