package main

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
