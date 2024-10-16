package db

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            int    `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBName            string `mapstructure:"DB_NAME"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBMaxOpenConns    int    `mapstructure:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns    int    `mapstructure:"DB_MAX_IDLE_CONNS"`
	DBConnMaxLifetime int    `mapstructure:"DB_CONN_MAX_LIFETIME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.SetDefault("DB_MAX_OPEN_CONNS", 10)
	viper.SetDefault("DB_MAX_IDLE_CONNS", 5)
	viper.SetDefault("DB_CONN_MAX_LIFETIME", 5)

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found, using environment variables only")
		} else {
			return
		}
	}

	err = viper.Unmarshal(&config)
	return
}

func JWTSecret() string {
	return viper.GetString("JWT_SECRET")
}

func ServerAddr() string {
	return viper.GetString("SERVER_ADDR")
}
