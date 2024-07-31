package types

type GetDataRequest struct{}

type GetDataResponse struct {
	Data string
}

type PingRequest struct{}

type PingResponse struct {
	Message string
}
