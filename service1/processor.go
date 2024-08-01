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

func NewRequestProcessor(min, max int, serviceName string) *RequestProcessor {
	return &RequestProcessor{
		minDelay:    min,
		maxDelay:    max,
		serviceName: serviceName,
	}
}

func (rp *RequestProcessor) Process(req *types.GetDataRequest, res *types.GetDataResponse) error {
	var delay time.Duration

	// make sure the maxDelay is greater than zero, otherwise the math.Intn will panic
	if rp.maxDelay > 0 {
		delay = time.Duration(rp.minDelay+rand.Intn(rp.getMaxDelay())) * time.Millisecond
	} else {
		delay = time.Duration(rp.minDelay) * time.Millisecond
	}

	time.Sleep(delay)

	res.Data = "Data from " + rp.serviceName

	return nil
}

func (rp *RequestProcessor) getMaxDelay() int {
	if rp.minDelay < 0 || rp.maxDelay < 0 {
		return 0
	}
	if rp.maxDelay > rp.minDelay {
		return rp.maxDelay - rp.minDelay
	}
	return 0
}
