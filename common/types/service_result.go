package types

import "time"

type ResultStatus int

const (
	ResultStatusSuccess ResultStatus = iota
	ResultStatusFailure
)

type ServiceResultEntry struct {
	ID          uint64
	DateTime    time.Time
	Status      ResultStatus
	Svc1Latency int64
	Svc2Latency int64
}
