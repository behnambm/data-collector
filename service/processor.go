package main

import (
	"github.com/behnambm/data-collector/common/types"
	"math/rand"
	"time"
)

type RequestProcessor struct {
	minDelay    int
	maxDelay    int
	serviceName string
}

func NewRequestProcessor(cfg *Config) (*RequestProcessor, error) {
	return &RequestProcessor{
		minDelay:    cfg.MinDelay,
		maxDelay:    cfg.GetMaxDelay(),
		serviceName: cfg.ServiceName,
	}, nil
}

func (rp *RequestProcessor) Process(req *types.GetDataRequest, res *types.GetDataResponse) error {
	var delay time.Duration

	// make sure the maxDelay is greater than zero, otherwise the math.Intn will panic
	if rp.maxDelay > 0 {
		delay = time.Duration(rp.minDelay+rand.Intn(rp.maxDelay)) * time.Millisecond
	} else {
		delay = time.Duration(rp.minDelay) * time.Millisecond
	}

	time.Sleep(delay)

	res.Data = "Data from " + rp.serviceName

	return nil
}
