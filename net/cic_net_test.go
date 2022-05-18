package net

import (
	"os"
	"testing"

	"github.com/grassrootseconomics/cic-go/provider"
	"github.com/lmittmann/w3"
)

type tConfig struct {
	rpcProvider string
	tokenIndex  string
	privateKey  string
}

var conf = &tConfig{
	rpcProvider: os.Getenv("RPC_PROVIDER"),
	tokenIndex:  os.Getenv("TOKEN_INDEX"),
	privateKey:  os.Getenv("PRIVATE_KEY"),
}

func TestCicNet_Connect(t *testing.T) {
	name := "Test CicNet connect"
	wantErr := false

	newProvider, err := provider.NewRpcProvider(conf.rpcProvider)
	if err != nil {
		t.Errorf("Creating an rpc instance failed = %v", err)
	}

	cicnet, _ := NewCicNet(*newProvider, w3.A(conf.tokenIndex))

	t.Run(name, func(t *testing.T) {
		if err := cicnet.provider.EthClient.Close(); (err != nil) != wantErr {
			t.Errorf("Error() error = %v, wantErr %v", err, wantErr)
		}
	})
}
