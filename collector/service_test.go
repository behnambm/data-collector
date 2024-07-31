package main

import (
	"errors"
	dummyDialler "github.com/behnambm/data-collector/collector/dialler/mock"
	dummyDB "github.com/behnambm/data-collector/collector/storage/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewService(t *testing.T) {
	db := new(dummyDB.DummyDB)
	dialler := new(dummyDialler.DummyDialler)
	cfg := &Config{}

	dialler.On("SetupTarget", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil)

	service, err := NewService(cfg, dialler, db)
	assert.NoError(t, err)
	assert.NotNil(t, service)
	assert.NotEmpty(t, service)
}

func TestNewService_ErrDialler_SetupTarget(t *testing.T) {
	db := new(dummyDB.DummyDB)
	dialler := new(dummyDialler.DummyDialler)
	cfg := &Config{}

	dialler.On("SetupTarget", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(errors.New("setup failed")).Times(2)

	service, err := NewService(cfg, dialler, db)

	assert.Error(t, err)
	assert.Nil(t, service)
}

func TestPing(t *testing.T) {
	db := new(dummyDB.DummyDB)
	dialler := new(dummyDialler.DummyDialler)
	cfg := &Config{}
	service := Service{cfg: cfg, dialler: dialler, storage: db}

	dialler.On(
		"Call",
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("*types.PingRequest"),
		mock.AnythingOfType("*types.PingResponse"),
	).
		Return(nil)

	err := service.Ping()

	assert.NoError(t, err)
}

func TestPing_ErrDialler_Call(t *testing.T) {
	db := new(dummyDB.DummyDB)
	dialler := new(dummyDialler.DummyDialler)
	cfg := &Config{}
	service := Service{cfg: cfg, dialler: dialler, storage: db}

	dialler.On(
		"Call",
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("*types.PingRequest"),
		mock.AnythingOfType("*types.PingResponse"),
	).
		Return(errors.New("ping failed")).Times(2)

	err := service.Ping()

	assert.Error(t, err)
}

func TestGetResult(t *testing.T) {
	db := new(dummyDB.DummyDB)
	dialler := new(dummyDialler.DummyDialler)
	dialler.Delay = 50
	cfg := &Config{
		Service1Name: "svc1",
		Service2Name: "svc2",
		Timeout:      500,
	}

	dialler.On("SetupTarget", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil)
	dialler.On(
		"Call",
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.Anything,
		mock.Anything,
	).Return(nil)

	db.On("Store", mock.Anything, mock.AnythingOfType("*types.ServiceResultEntry")).
		Return(nil)

	service, err := NewService(cfg, dialler, db)
	err = service.GetResult()

	assert.NoError(t, err)
}

func TestGetResult_ErrTimeout(t *testing.T) {
	db := new(dummyDB.DummyDB)
	dialler := new(dummyDialler.DummyDialler)
	dialler.Delay = 60
	cfg := &Config{
		Service1Name: "svc1",
		Service2Name: "svc2",
		Timeout:      50,
	}

	dialler.On("SetupTarget", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil)
	dialler.On(
		"Call",
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.Anything,
		mock.Anything,
	).Return(nil)

	db.On("Store", mock.Anything, mock.AnythingOfType("*types.ServiceResultEntry")).
		Return(nil)

	service, err := NewService(cfg, dialler, db)
	err = service.GetResult()

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrTimeout)
}

func TestGetResult_ErrStoreResultFailed(t *testing.T) {
	db := new(dummyDB.DummyDB)
	dialler := new(dummyDialler.DummyDialler)
	dialler.Delay = 50
	cfg := &Config{
		Service1Name: "svc1",
		Service2Name: "svc2",
		Timeout:      500,
	}

	dialler.On("SetupTarget", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil)
	dialler.On(
		"Call",
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.Anything,
		mock.Anything,
	).Return(nil)

	db.On("Store", mock.Anything, mock.AnythingOfType("*types.ServiceResultEntry")).
		Return(errors.New("storage failed"))

	service, err := NewService(cfg, dialler, db)
	err = service.GetResult()

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrStoreResultFailed)
}
