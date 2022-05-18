package balance

import (
	"os"
	"testing"

	"github.com/grassrootseconomics/cic-go/provider"
	"github.com/lmittmann/w3"
)

type tConfig struct {
	rpcProvider   string
	batchContract string
}

var conf = &tConfig{
	rpcProvider:   os.Getenv("RPC_PROVIDER"),
	batchContract: os.Getenv("BATCH_CONTRACT"),
}

func TestBatchBalance_Connect(t *testing.T) {
	name := "Test RPC connection"
	wantErr := false

	newProvider, err := provider.NewRpcProvider(conf.rpcProvider)
	if err != nil {
		t.Errorf("Creating an rpc instance failed = %v", err)
	}

	batchBalance, _ := NewBatchBalance(*newProvider, w3.A(conf.batchContract))

	t.Run(name, func(t *testing.T) {
		if err := batchBalance.provider.EthClient.Close(); (err != nil) != wantErr {
			t.Errorf("Close() error = %v, wantErr %v", err, wantErr)
		}
	})
}
