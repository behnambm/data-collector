package rpc

import (
	"errors"
	"net/rpc"
)

var (
	ErrClientDoesNotExist = errors.New("RPC client does not exist")
)

type RPCDialler struct {
	// when a target service gets setup, the client to that RPC server is stored in clientStore
	clientStore map[string]*rpc.Client
}

func New() (*RPCDialler, error) {
	return &RPCDialler{
		clientStore: make(map[string]*rpc.Client),
	}, nil
}

// SetupTarget will create a client to the target service and store the client(connection) for future re-usability
// this way the handshake overhead will be lower
func (rd *RPCDialler) SetupTarget(targetServiceName, serviceAddress string) error {
	client, err := rpc.Dial("tcp", serviceAddress)
	if err != nil {
		return err
	}
	rd.clientStore[targetServiceName] = client
	return nil
}

func (rd *RPCDialler) getClient(serviceName string) (*rpc.Client, error) {
	if client, ok := rd.clientStore[serviceName]; ok {
		return client, nil
	}
	return nil, ErrClientDoesNotExist
}

func (rd *RPCDialler) Call(targetServiceName, methodName string, req, res any) error {
	client, err := rd.getClient(targetServiceName)
	if err != nil {
		return err
	}

	if err = client.Call(methodName, req, res); err != nil {
		return err
	}
	return nil
}

func (rd *RPCDialler) Close() {
	for _, client := range rd.clientStore {
		client.Close()
	}
}
