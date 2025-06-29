package configs

import (
	"strings"
)

type KafkaProducerConfig struct {
	HostsString string `env:"HOSTS,required"`
	Topic       string `env:"TOPIC,required"`
}

func (s KafkaProducerConfig) Hosts() []string {
	return strings.Split(s.HostsString, ",")
}

type KafkaConsumerConfig struct {
	HostsString   string `env:"HOSTS,required"`
	Topic         string `env:"TOPIC,required"`
	ConsumerGroup string `env:"CONSUMER_GROUP,required"`
}

func (s KafkaConsumerConfig) Hosts() []string {
	return strings.Split(s.HostsString, ",")
}
