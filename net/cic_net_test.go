package net

import (
	"github.com/lmittmann/w3"
	"os"
	"testing"
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
	name := "Test RPC connection"
	wantErr := false

	cicnet, _ := NewCicNet(conf.rpcProvider, w3.A(conf.tokenIndex))

	t.Run(name, func(t *testing.T) {
		if err := cicnet.Close(); (err != nil) != wantErr {
			t.Errorf("EntryCount() error = %v, wantErr %v", err, wantErr)
		}
	})
}
