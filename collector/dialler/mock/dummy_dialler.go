package mock

import (
	"github.com/stretchr/testify/mock"
	"time"
)

type DummyDialler struct {
	mock.Mock
	Delay int
}

func (d DummyDialler) SetupTarget(targetServiceName, serviceAddress string) error {
	args := d.Called(targetServiceName, serviceAddress)
	return args.Error(0)
}

func (d DummyDialler) Call(targetServiceName, methodName string, req, res any) error {
	args := d.Called(targetServiceName, methodName, req, res)
	time.Sleep(time.Duration(d.Delay) * time.Millisecond)
	return args.Error(0)
}
