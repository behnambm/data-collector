package main

import (
	"fmt"
	"github.com/behnambm/data-collector/config"
	log "github.com/sirupsen/logrus"
	"net"
	"net/rpc"
)

func main() {
	configFilePath := config.ParseArgs()
	cfg, err := config.LoadConfig[Config](configFilePath)
	if err != nil {
		log.Fatalf("Unable to load configs: %v\n", err)
	}

	reqProcessor, err := NewRequestProcessor(cfg)
	if err != nil {
		log.Fatalf("Unable to initialize request processor: %v\n", err)
	}

	svc1RPC, err := NewService1RPC(reqProcessor)
	if err != nil {
		log.Fatalf("Unable to initialize service: %v\n", err)
	}

	server := rpc.NewServer()
	err = server.RegisterName("ServiceRPC", svc1RPC)
	if err != nil {
		log.Fatalf("Unable to initialize RPC: %v\n", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatalf("Unable to bind RPC server to %s:%d error: %v\n", cfg.Host, cfg.Port, err)
	}

	log.WithField("service name", cfg.ServiceName).Infof(
		"[%s] started listening on RPC... (%s:%d)",
		cfg.ServiceName, cfg.Host, cfg.Port,
	)

	server.Accept(listener)
}
