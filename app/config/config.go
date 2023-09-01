package config

import (
	"os"
	"time"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
)

func GetKafkaConfig() *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	return kafkaConfig
}

func Env(key string) string {
	_ = godotenv.Load(".env")
	return os.Getenv(key)
}
