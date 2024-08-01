package main

import (
	"github.com/behnambm/data-collector/common/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		name        string
		min         int
		max         int
		expectedMin int
		expectedMax int
	}{
		{
			name:        "Basic test",
			min:         10,
			max:         20,
			expectedMin: 10,
			expectedMax: 20,
		},
		{
			name:        "Zero max",
			min:         10,
			max:         0,
			expectedMin: 10,
			expectedMax: 10,
		},
		{
			name:        "Zero min",
			min:         0,
			max:         10,
			expectedMin: 0,
			expectedMax: 10,
		},
		{
			name:        "Negative delays",
			min:         -5,
			max:         -10,
			expectedMin: 0,
			expectedMax: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rp := &RequestProcessor{
				minDelay: test.min,
				maxDelay: test.max,
			}

			req := &types.GetDataRequest{}
			res := &types.GetDataResponse{}

			start := time.Now()
			rp.Process(req, res)
			elapsed := int(time.Since(start).Milliseconds())

			assert.LessOrEqual(t, elapsed, test.expectedMax)
			assert.GreaterOrEqual(t, elapsed, test.expectedMin)
		})
	}
}

func TestConfig_GetMaxDelay(t *testing.T) {
	tests := []struct {
		name     string
		min      int
		max      int
		expected int
	}{
		{
			"Basic test",
			300,
			600,
			300,
		},
		{
			"Negative min",
			-1,
			5,
			0,
		},
		{
			"Negative max",
			5,
			-1,
			0,
		},
		{
			"min greater than max",
			10,
			5,
			0,
		},
		{
			"min equal to max",
			5,
			5,
			0,
		},
		{
			"max greater than min",
			5,
			10,
			5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rp := &RequestProcessor{
				minDelay: test.min,
				maxDelay: test.max,
			}
			result := rp.getMaxDelay()
			if result != test.expected {
				assert.Equal(t, test.expected, result)
			}
		})
	}
}
