package main

import (
	"github.com/behnambm/data-collector/common/types"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1111")
	if err != nil {
		log.Fatalf("Error dialing RPC server: %v", err)
	}

	req := &types.PingRequest{}
	res := &types.PingResponse{}

	err = client.Call("ServiceRPC.Ping", req, res)
	if err != nil {
		log.Fatalf("Error calling Ping method: %v", err)
	}

	log.Printf("Ping response: %s", res.Message)
}
