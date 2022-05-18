package net

import (
	"context"
	"math/big"
	"testing"

	"github.com/grassrootseconomics/cic-go/provider"
	"github.com/lmittmann/w3"
)

func TestCicNet_TokenIndex_EntryCount(t *testing.T) {
	name := "Entry count"
	wantErr := false

	newProvider, err := provider.NewRpcProvider(conf.rpcProvider)
	if err != nil {
		t.Errorf("Creating an rpc instance failed = %v", err)
	}
	cicnet, err := NewCicNet(newProvider, w3.A(conf.tokenIndex))

	if err != nil {
		t.Fatalf("NewCicNet error = %v", err)
	}

	t.Run(name, func(t *testing.T) {
		got, err := cicnet.EntryCount(context.Background())

		if (err != nil) != wantErr {
			t.Errorf("EntryCount() error = %v, wantErr %v", err, wantErr)
		}

		if got.Cmp(big.NewInt(0)) < 1 {
			t.Errorf("EntryCount() = %v, want %v", got, 1)
		}
	})
}

func TestCicNet_TokenIndex_AddressAtIndex(t *testing.T) {
	type args struct {
		index *big.Int
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		address string
	}{
		{
			name: "Address at index 0",
			args: args{
				index: big.NewInt(0),
			},
			wantErr: false,
			address: "0xaB89822F31c2092861F713F6F34bd6877a8C1878",
		},
		{
			name: "Address at index 999",
			args: args{
				index: big.NewInt(999),
			},
			wantErr: true,
			address: "",
		},
	}

	newProvider, err := provider.NewRpcProvider(conf.rpcProvider)
	if err != nil {
		t.Errorf("Creating an rpc instance failed = %v", err)
	}
	cicnet, err := NewCicNet(newProvider, w3.A(conf.tokenIndex))

	if err != nil {
		t.Fatalf("NewCicNet error = %v", err)
	}

	for _, testcase := range tests {
		tt := testcase

		t.Run(tt.name, func(t *testing.T) {
			got, err := cicnet.AddressAtIndex(context.Background(), tt.args.index)

			if (err != nil) != tt.wantErr {
				t.Errorf("AddressAtIndex() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.address {
				t.Errorf("AddressAtIndex = %v, want %v", got, tt.address)
			}
		})
	}
}
