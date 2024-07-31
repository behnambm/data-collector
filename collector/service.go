package main

import (
	"context"
	"fmt"
	"github.com/behnambm/data-collector/common/types"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Dialler interface {
	SetupTarget(targetServiceName, serviceAddress string) error
	Call(targetServiceName, methodName string, req, res any) error
}

type Storage interface {
	Store(ctx context.Context, entry *types.ServiceResultEntry) error
}

type Service struct {
	dialler Dialler
	storage Storage
	cfg     *Config
	doneCh  chan struct{}
	result  map[string]int64
	wg      sync.WaitGroup
}

func NewService(cfg *Config, dialler Dialler, storage Storage) (*Service, error) {
	svc1Address := fmt.Sprintf("%s:%d", cfg.Service1Host, cfg.Service1Port)
	svc2Address := fmt.Sprintf("%s:%d", cfg.Service2Host, cfg.Service2Port)

	if err := dialler.SetupTarget(cfg.Service1Name, svc1Address); err != nil {
		return nil, err
	}
	if err := dialler.SetupTarget(cfg.Service2Name, svc2Address); err != nil {
		return nil, err
	}

	return &Service{
		dialler: dialler,
		storage: storage,
		cfg:     cfg,
		doneCh:  make(chan struct{}),
		result:  make(map[string]int64),
		wg:      sync.WaitGroup{},
	}, nil
}

func (s *Service) Ping() error {
	err := s.dialler.Call(
		s.cfg.Service1Name,
		"ServiceRPC.Ping",
		&types.PingRequest{},
		&types.PingResponse{},
	)
	if err != nil {
		return err
	}

	err = s.dialler.Call(
		s.cfg.Service2Name,
		"ServiceRPC.Ping",
		&types.PingRequest{},
		&types.PingResponse{},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetResult() error {
	s.wg.Add(2)
	go s.getData(s.cfg.Service1Name)
	go s.getData(s.cfg.Service2Name)
	go func() { s.wg.Wait(); s.doneCh <- struct{}{} }()

	select {
	case <-s.doneCh:
		log.Infoln("ALL DONE")
		return s.storeResult(types.ResultStatusSuccess)
	case <-time.After(time.Duration(s.cfg.Timeout) * time.Millisecond):
		log.Errorln("TIMEOUT")
		return s.storeResult(types.ResultStatusFailure)
	}
}

func (s *Service) getData(serviceName string) {
	defer s.wg.Done()

	req := &types.GetDataRequest{}
	res := &types.GetDataResponse{}

	start := time.Now()

	// TODO: the returning error of Call also can be stored in ServiceResult to be able to process afterward
	s.dialler.Call(serviceName, "ServiceRPC.GetData", req, res)
	elapsed := time.Since(start)

	fmt.Printf("time elapsed for (%s): %d\n", serviceName, elapsed.Milliseconds())

	s.result[serviceName] = elapsed.Milliseconds()
}

func (s *Service) storeResult(status types.ResultStatus) error {
	svc1Latency, ok := s.result[s.cfg.Service1Name]
	if !ok {
		// setting to empty struct because we also want to save the failed requests
		// if we don't need failed ones then we can simply return an error instead
		svc1Latency = 0
	}

	svc2Latency, ok := s.result[s.cfg.Service2Name]
	if !ok {
		svc2Latency = 0
	}

	entry := types.ServiceResultEntry{
		DateTime:    time.Now(),
		Status:      status,
		Svc1Latency: svc1Latency,
		Svc2Latency: svc2Latency,
	}

	if err := s.storage.Store(context.TODO(), &entry); err != nil {
		return err
	}

	return nil
}
