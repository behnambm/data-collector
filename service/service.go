package main

import (
	"github.com/behnambm/data-collector/common/types"
	log "github.com/sirupsen/logrus"
)

type ServiceRPC struct {
	reqProcessor *RequestProcessor
}

func NewServiceRPC(reqProcessor *RequestProcessor) (*ServiceRPC, error) {
	return &ServiceRPC{
		reqProcessor: reqProcessor,
	}, nil
}

func (s *ServiceRPC) GetData(req *types.GetDataRequest, res *types.GetDataResponse) error {
	log.WithField("method", "GetData").Info("handling request...")
	return s.reqProcessor.Process(req, res)
}

func (s *ServiceRPC) Ping(req *types.PingRequest, res *types.PingResponse) error {
	log.WithField("method", "Ping").Info("handling request...")
	// check some stuff and make sure the service is stable
	res.Message = "PONG"

	return nil
}
