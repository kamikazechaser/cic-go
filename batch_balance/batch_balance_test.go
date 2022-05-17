package balance

import (
	"os"
	"testing"

	"github.com/lmittmann/w3"
)

type tConfig struct {
	rpcProvider   string
	batchContract string
}

var conf = &tConfig{
	rpcProvider:   os.Getenv("RPC_PROVIDER"),
	batchContract: os.Getenv("BATCH_BALANCE"),
}

func TestBatchBalance_Connect(t *testing.T) {
	name := "Test RPC connection"
	wantErr := false

	cicnet, _ := NewBatchBalance(conf.rpcProvider, w3.A(conf.batchContract))

	t.Run(name, func(t *testing.T) {
		if err := cicnet.Close(); (err != nil) != wantErr {
			t.Errorf("Close() error = %v, wantErr %v", err, wantErr)
		}
	})
}
