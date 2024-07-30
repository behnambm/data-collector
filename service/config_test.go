package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_GetMaxDelay(t *testing.T) {
	tests := []struct {
		name     string
		config   Config
		expected int
	}{
		{
			name:     "Basic case",
			config:   Config{MinDelay: 100, MaxDelay: 200},
			expected: 100,
		},
		{
			name:     "Zero delays",
			config:   Config{MinDelay: 0, MaxDelay: 0},
			expected: 0,
		},
		{
			name:     "Negative MinDelay",
			config:   Config{MinDelay: -100, MaxDelay: 200},
			expected: 0,
		},
		{
			name:     "Negative MaxDelay",
			config:   Config{MinDelay: 100, MaxDelay: -200},
			expected: 0,
		},
		{
			name:     "Both negative delays",
			config:   Config{MinDelay: -100, MaxDelay: -200},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.config.GetMaxDelay()
			assert.Equal(t, test.expected, result, "GetMaxDelay failed for case: %s", test.name)
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		config   Config
		expected error
	}{
		{
			name:     "Valid Config",
			config:   Config{MinDelay: 100, MaxDelay: 200},
			expected: nil,
		},
		{
			name:     "Both Delays zero",
			config:   Config{MinDelay: 0, MaxDelay: 0},
			expected: nil,
		},
		{
			name:     "Negative MinDelay",
			config:   Config{MinDelay: -1, MaxDelay: 200},
			expected: ErrMinDelay,
		},
		{
			name:     "Negative MaxDelay",
			config:   Config{MinDelay: 100, MaxDelay: -1},
			expected: ErrMaxDelay,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.config.Validate()
			assert.ErrorIs(t, test.expected, err, "GetMaxDelay failed for case: %s", test.name)
		})
	}
}
