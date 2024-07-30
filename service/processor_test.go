package main

import (
	"github.com/behnambm/data-collector/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestProcess_SleepsInDefinedRange(t *testing.T) {
	cfg := &Config{
		MinDelay:    100,
		MaxDelay:    200,
		ServiceName: "test",
	}

	rp, err := NewRequestProcessor(cfg)
	assert.NoError(t, err, "failed to initialize RequestProcessor")

	req := &types.GetDataRequest{}
	res := &types.GetDataResponse{}

	startTime := time.Now()
	err = rp.Process(req, res)
	endTime := time.Now()

	assert.NoError(t, err, "failed to process request")

	elapsedTime := endTime.Sub(startTime)
	minExpected := time.Duration(cfg.MinDelay) * time.Millisecond
	maxExpected := time.Duration(cfg.MinDelay+cfg.MaxDelay) * time.Millisecond

	assert.GreaterOrEqual(t, elapsedTime, minExpected, "process duration is less than expected")
	assert.LessOrEqual(t, elapsedTime, maxExpected, "process duration is greater than expected")
	assert.Equal(t, "Data from "+cfg.ServiceName, res.Data, "unexpected response data")
}

func TestProcess_ZeroDelay(t *testing.T) {
	cfg := &Config{
		MinDelay:    0,
		MaxDelay:    0,
		ServiceName: "test",
	}

	rp, err := NewRequestProcessor(cfg)
	assert.NoError(t, err, "failed to create RequestProcessor")

	req := &types.GetDataRequest{}
	res := &types.GetDataResponse{}

	start := time.Now()
	err = rp.Process(req, res)
	end := time.Now()

	assert.NoError(t, err, "failed to process request")

	elapsedTime := end.Sub(start)

	// setting expected to 1 second because the code itself will take more than zero second ro tun
	expected := 1 * time.Millisecond

	assert.LessOrEqual(t, elapsedTime, expected, "process duration is greater than expected")
	assert.Equal(t, "Data from "+cfg.ServiceName, res.Data, "unexpected response data")
}
