package main

import (
	"fmt"
	"github.com/behnambm/data-collector/common/config"
	log "github.com/sirupsen/logrus"
	"net"
	"net/rpc"
	"os"
	"os/signal"
)

func main() {
	configFilePath := config.ParseArgs()
	cfg, err := config.LoadConfig[Config](configFilePath)
	if err != nil {
		log.Fatalf("Unable to load configs: %v\n", err)
	}

	reqProcessor := NewRequestProcessor(cfg.MinDelay, cfg.MaxDelay, cfg.ServiceName)

	svcRPC, err := NewServiceRPC(reqProcessor)
	if err != nil {
		log.Fatalf("Unable to initialize service: %v\n", err)
	}

	server := rpc.NewServer()
	err = server.RegisterName("ServiceRPC", svcRPC)
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

	go server.Accept(listener)

	// handle graceful shutdown
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Kill, os.Interrupt)
	sig := <-signalCh

	// this will make OS be able to do force close
	signal.Reset(sig)
	log.Infoln("Shutting down, please wait...")
	listener.Close()
}
