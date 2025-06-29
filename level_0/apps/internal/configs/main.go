package configs

import (
	"github.com/caarlos0/env/v11"
	"log"
)

type WebAndConsumerConfig struct {
	KafkaConsumer KafkaConsumerConfig `envPrefix:"KAFKA__"`
	Postgres      PostgresConfig      `envPrefix:"POSTGRES__"`
}

type ProducerConfig struct {
	KafkaProducer KafkaProducerConfig `envPrefix:"KAFKA__"`
}

type allowedConfig interface {
	WebAndConsumerConfig | ProducerConfig
}

// Получение конфига
func GetConfig[CT allowedConfig]() *CT {
	var conf CT
	if err := env.Parse(&conf); err != nil {
		log.Fatal(err)
	}
	return &conf
}
