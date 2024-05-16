package config

import "os"

type Config struct {
	KafkaBroker string
}

func LoadConfig() Config {
	return Config{
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
	}
}
