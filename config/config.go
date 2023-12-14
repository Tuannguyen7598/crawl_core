package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port    int
	BaseUrl string
}

func GetConfigAvaiable() Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000
	}
	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "localhost:3000"
	}

	var Config = Config{
		Port:    port,
		BaseUrl: baseUrl,
	}
	return Config

}
