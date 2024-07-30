package main

import "errors"

var (
	ErrMinDelay = errors.New("MinDelay cannot be negative")
	ErrMaxDelay = errors.New("MaxDelay cannot be negative")
)

type Config struct {
	// Service specific configs
	ServiceName string

	// both MinDelay and MaxDelay are in milliseconds
	MinDelay int
	MaxDelay int

	// RPC specific configs
	Host string
	Port int
}

func (c *Config) Validate() error {
	if c.MinDelay < 0 {
		return ErrMinDelay
	}
	if c.MaxDelay < 0 {
		return ErrMaxDelay
	}

	return nil
}

func (c *Config) GetMaxDelay() int {
	if c.MinDelay < 0 || c.MaxDelay < 0 {
		return 0
	}
	if c.MaxDelay > c.MinDelay {
		return c.MaxDelay - c.MinDelay
	}
	return 0
}
