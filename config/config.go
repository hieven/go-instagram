package config

import "time"

type Config struct {
	Username      string
	Password      string
	Capacity      int
	LoginInterval time.Duration
}
