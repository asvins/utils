package config

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

// Load a config from file and return a Config object
func Load(filepath string, cfg interface{}) error {
	err := gcfg.ReadFileInto(cfg, filepath)
	if err != nil {
		fmt.Println("Error while loading config \n", err)
		return err
	}
	return nil
}
