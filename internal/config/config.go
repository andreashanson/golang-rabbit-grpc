package config

import "os"

type Config struct {
	RabbitConfig *RabbitConfig
}

type RabbitConfig struct {
	AmqpServerUrl string
	Queue         string
}

func NewConfig() *Config {
	return &Config{RabbitConfig: &RabbitConfig{AmqpServerUrl: os.Getenv("AMQP_SERVER_URL"), Queue: os.Getenv("RABBIT_QUEUE")}}
}
