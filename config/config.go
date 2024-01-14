package config

import (
	"os"
)

type Config struct {
	Port              string
	DatabaseHost      string    
	DatabaseName      string 
	DatabaseUser      string 
	DatabasePassword  string 
	DatabasePort      string 
	DatabaseSSLMODE   string 
}

func GetConfig() *Config {
	return &Config{
		Port              :os.Getenv("APP_CONTAINER_PORT"),  
		DatabaseHost      :os.Getenv("POSTGRES_HOST"),
		DatabaseName      :os.Getenv("POSTGRES_DB"),  
		DatabaseUser      :os.Getenv("POSTGRES_USER"),  
		DatabasePassword  :os.Getenv("POSTGRES_PASSWORD"),   
		DatabasePort      :os.Getenv("POSTGRES_CONTAINER_PORT"),   
		DatabaseSSLMODE   :os.Getenv("POSTGRES_SSLMODE"),   
	}
}
