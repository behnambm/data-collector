package main

import (
	"fmt"
	"github.com/behnambm/data-collector/collector/dialler/rpc"
	"github.com/behnambm/data-collector/collector/storage/sqlite"
	"github.com/behnambm/data-collector/common/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	configFilePath := config.ParseArgs()
	cfg, err := config.LoadConfig[Config](configFilePath)
	if err != nil {
		log.Fatalf("Unable to load configs: %v\n", err)
	}

	sqliteStorage, err := sqlite.New(&cfg.DBConfig)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	if err = sqliteStorage.SetupModels(); err != nil {
		log.Fatalf("Unable to setup database: %v\n", err)
	}

	rpcDialler, err := rpc.New()
	if err != nil {
		log.Fatalf("Unable to initialize RPC dialler: %v\n", err)
	}

	service, err := NewService(cfg, rpcDialler, sqliteStorage)
	if err != nil {
		log.Fatalf("Unable to initialize service: %v\n", err)
	}
	if err = service.Ping(); err != nil {
		log.Fatalf("Unable to ping services: %v\n", err)
	}

	if err = service.GetResult(); err != nil {
		fmt.Println("GetResult error: ", err)
	}
}
