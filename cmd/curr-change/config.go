package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type configuration struct {
	Addr string `env:"ADDR" default:"127.0.0.1"`
	Port string `env:"Port" default:"8888"`

	DBConfig string `env:"DB"`

	LogLevel int `env:"VERBOSE" default:"1"`
}

func parseConfig() (*configuration, error) {
	c := &configuration{}

	if err := env.Parse(c); err != nil {
		return nil, fmt.Errorf("parsing envs: %w", err)
	}

	return c, nil
}
