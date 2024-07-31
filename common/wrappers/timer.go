package wrappers

import "time"

func Timer(f func()) func() time.Duration {
	return func() time.Duration {
		start := time.Now()
		f()
		elapsed := time.Since(start)
		return elapsed
	}
}
