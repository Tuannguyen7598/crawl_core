package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port    int
	BaseUrl string
	Uri_mongo string
	DB_Mongo string
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
		Uri_mongo: "mongodb://user:pass@localhost:27017",
		DB_Mongo: "crawl_data",
	}
	return Config

}
