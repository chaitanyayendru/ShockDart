package config

import (
    "os"
)

type Config struct {
    ServerPort        string
    KafkaBroker       string
    FirebaseProjectID string
}

func LoadConfig() Config {
    return Config{
        ServerPort:        getEnv("SERVER_PORT", ":8080"),
        KafkaBroker:       getEnv("KAFKA_BROKER", "localhost:9092"),
        FirebaseProjectID: getEnv("FIREBASE_PROJECT_ID", ""),
    }
}

func getEnv(key, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultVal
}
