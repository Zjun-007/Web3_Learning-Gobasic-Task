package config

import (
    "os"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    JWTSecret  string
    JWTExpire  int
    ServerPort string
}

var AppConfig *Config   

func LoadConfig() *Config {
    AppConfig = &Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "3306"),
        DBUser:     getEnv("DB_USER", "godb"),
        DBPassword: getEnv("DB_PASSWORD", "54862"),
        DBName:     getEnv("DB_NAME", "blogdb"),
        JWTSecret:  getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
        ServerPort: getEnv("SERVER_PORT", "8080"),
    }
    return AppConfig    
}   

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}