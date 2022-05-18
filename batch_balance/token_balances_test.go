package balance

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/grassrootseconomics/cic-go/provider"
	"github.com/lmittmann/w3"
)

func TestBatchBalance_TokensBalance(t *testing.T) {
	type args struct {
		owner  common.Address
		tokens []common.Address
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "A random members (min dust available) balances",
			args: args{
				owner: w3.A("0x4e956b5De3c33566c596754B4fa0ABd9F2789578"),
				tokens: []common.Address{
					w3.A("0xaB89822F31c2092861F713F6F34bd6877a8C1878"),
					w3.A("0x982caeF20362ADEAC3f9a25E37E20E6787f27f85"),
					w3.A("0x9ADd261033baA414c84FF84A4Fe396338C1ba13a"),
					w3.A("0x7dF20b526318d37Cd7DA9518E51d4A51fec30c9A"),
				},
			},
			want:    true,
			wantErr: false,
		},
	}

	newProvider, err := provider.NewRpcProvider(conf.rpcProvider)
	if err != nil {
		t.Errorf("Creating an rpc instance failed = %v", err)
	}
	batchBalance, err := NewBatchBalance(newProvider, w3.A(conf.batchContract))

	if err != nil {
		t.Fatalf("NewBatchBalance error = %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := batchBalance.TokensBalance(context.Background(), tt.args.owner, tt.args.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchBalance.TokensBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got[2].Int64() > 0 {
				t.Errorf("TokenBalance = %d, want %d", got[2].Int64(), 0)
			}

		})
	}
}
