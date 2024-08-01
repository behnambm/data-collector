package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/behnambm/data-collector/common/types"
	"github.com/behnambm/data-collector/common/wrappers"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	ErrTimeout           = errors.New("timeout")
	ErrStoreResultFailed = errors.New("storing result failed")
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
	result  sync.Map
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
		result:  sync.Map{},
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
		if err := s.storeResult(types.ResultStatusSuccess); err != nil {
			log.Errorln("storeResult error:", err)
			return ErrStoreResultFailed
		}
		return nil
	case <-time.After(time.Duration(s.cfg.Timeout) * time.Millisecond):
		log.Errorln("TIMEOUT")
		if err := s.storeResult(types.ResultStatusFailure); err != nil {
			log.Errorln("storeResult error:", err)
			return ErrStoreResultFailed
		}
		return ErrTimeout
	}
}

func (s *Service) getData(serviceName string) {
	defer s.wg.Done()

	req := &types.GetDataRequest{}
	res := &types.GetDataResponse{}

	elapsed := wrappers.Timer(
		func() {
			// TODO: the returning error of Call also can be stored in ServiceResult to be able to process afterward
			s.dialler.Call(serviceName, "ServiceRPC.GetData", req, res)
		},
	)()

	log.Debugf("time elapsed for (%s): %d\n", serviceName, elapsed.Milliseconds())

	s.result.Store(serviceName, elapsed.Milliseconds())
}

func (s *Service) loadLatency(serviceName string) int64 {
	latency, ok := s.result.Load(serviceName)
	latencyInt64, ok := latency.(int64)
	if !ok {
		// returning zero because we also want to save the failed requests
		return 0
	}
	return latencyInt64
}

func (s *Service) storeResult(status types.ResultStatus) error {
	entry := types.ServiceResultEntry{
		DateTime:    time.Now(),
		Status:      status,
		Svc1Latency: s.loadLatency(s.cfg.Service1Name),
		Svc2Latency: s.loadLatency(s.cfg.Service2Name),
	}

	if err := s.storage.Store(context.TODO(), &entry); err != nil {
		return err
	}

	return nil
}
