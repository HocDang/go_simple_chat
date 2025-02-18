package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	PostgresHost           string `mapstructure:"POSTGRES_HOST"`
	PostgresPort           string `mapstructure:"POSTGRES_PORT"`
	PostgresUser           string `mapstructure:"POSTGRES_USER"`
	PostgresPass           string `mapstructure:"POSTGRES_PASS"`
	PostgresName           string `mapstructure:"POSTGRES_NAME"`
	RedisHost              string `mapstructure:"REDIS_HOST"`
	RedisPort              string `mapstructure:"REDIS_PORT"`
	RedisPass              string `mapstructure:"REDIS_PASS"`
	RedisName              int    `mapstructure:"REDIS_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func LoadConfig() *Env {
	env := Env{}
	viper.SetConfigFile("./.env")

	err := viper.ReadInConfig()
	if err != nil {
		currentDir, err := os.Getwd()
		log.Fatalf("Can't find the file .env in %s : %v", currentDir, err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
