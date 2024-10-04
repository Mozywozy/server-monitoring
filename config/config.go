package config

import "time"

type Config struct {
	Servers  []string
	Interval time.Duration
}

func LoadConfig() Config {
	return Config{
		Servers:  []string{"8.8.8.8", "1.1.1.1"},
		Interval: 10 * time.Second,
	}
}
