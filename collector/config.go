package main

import "github.com/behnambm/data-collector/collector/storage/sqlite"

type Config struct {
	// Service1Name is just for making a difference among other services
	Service1Name string
	Service1Host string
	Service1Port int

	// Service2Name is just for making a difference among other services
	Service2Name string
	Service2Host string
	Service2Port int

	// Timeout in milliseconds
	Timeout int

	DBConfig sqlite.Config
}
