package config

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

// Config struct
type Config struct {
	Server struct {
		Addr string
		Port string
	}
}

// Load a config from file and return a Config object
func Load(filepath string) (*Config, error) {
	cfg := Config{}
	err := gcfg.ReadFileInto(&cfg, filepath)
	if err != nil {
		fmt.Println("Error while loading config \n", err)
		return nil, err
	}
	return &cfg, nil
}
