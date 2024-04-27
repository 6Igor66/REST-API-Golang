package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Http       HTTPServer         `json:"http"`
	Postgresql postgresqlSettings `json:"postgresql"`
}

type postgresqlSettings struct {
	ConnString string `json:"connstring" required:"true"`
}

type HTTPServer struct {
	Port string `json:"port"`
}

func ReadConfigFromFile(filePath string) (*Config, error) {
	var config Config

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config data: %v", err)
		return nil, err
	}

	return &config, nil
}
