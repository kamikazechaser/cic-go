package cic_net

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"math/big"
	"testing"
)

func TestCicNet_DemurrageToken_DemurrageTokeInfo(t *testing.T) {
	type args struct {
		contractAddress common.Address
	}

	tests := []struct {
		name        string
		args        args
		wantErr     bool
		isDemurrage bool
	}{
		{
			name: "Demurrage token at kitabu sarafu",
			args: args{
				contractAddress: w3.A("0xaB89822F31c2092861F713F6F34bd6877a8C1878"),
			},
			wantErr: false,
		},
		{
			name: "Giftable token at Muthaa",
			args: args{
				contractAddress: w3.A("0x3dad47e5EF13661bbD15aa74132E91a9aBCFDe44"),
			},
			wantErr: true,
		},
		{
			name: "Dead address",
			args: args{
				contractAddress: w3.A("0x000000000000000000000000000000000000dEaD"),
			},
			wantErr: true,
		},
	}

	for _, testcase := range tests {
		tt := testcase

		t.Run(tt.name, func(t *testing.T) {
			cicnet, err := NewCicNet(conf.rpcProvider, w3.A(conf.tokenIndex))

			if err != nil {
				t.Fatalf("NewCicNet error = %v", err)
			}

			got, err := cicnet.DemurrageTokenInfo(context.Background(), tt.args.contractAddress)

			if (err != nil) != tt.wantErr {
				t.Errorf("DemurrageTokenInfo() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if got.DemurrageAmount.Cmp(big.NewInt(0)) < 1 {
					t.Fatalf("DemurrageAmount = %v, want %d atleast", got, 1)
				}
			}
		})
	}
}

func TestCicNet_DemurrageToken_BaseBalanceOf(t *testing.T) {
	type args struct {
		contractAddress common.Address
		accountAddress  common.Address
	}

	tests := []struct {
		name       string
		args       args
		wantErr    bool
		balanceGte big.Int
	}{
		{
			name: "Sarafu sink balance",
			args: args{
				contractAddress: w3.A("0xaB89822F31c2092861F713F6F34bd6877a8C1878"),
				accountAddress:  w3.A("0xBBb4a93c8dCd82465B73A143f00FeD4AF7492a27"),
			},
			wantErr:    false,
			balanceGte: *big.NewInt(1),
		},
		{
			name: "Dead address balance",
			args: args{
				contractAddress: w3.A("0xaB89822F31c2092861F713F6F34bd6877a8C1878"),
				accountAddress:  w3.A("0x000000000000000000000000000000000000dEaD"),
			},
			wantErr:    false,
			balanceGte: *big.NewInt(0),
		},
		{
			name: "Giftable token at Muthaa",
			args: args{
				contractAddress: w3.A("0x3dad47e5EF13661bbD15aa74132E91a9aBCFDe44"),
				accountAddress:  w3.A("0xBBb4a93c8dCd82465B73A143f00FeD4AF7492a27"),
			},
			wantErr: true,
		},
	}

	for _, testcase := range tests {
		tt := testcase

		t.Run(tt.name, func(t *testing.T) {
			cicnet, err := NewCicNet(conf.rpcProvider, w3.A(conf.tokenIndex))

			if err != nil {
				t.Fatalf("NewCicNet error = %v", err)
			}

			got, err := cicnet.BaseBalanceOf(context.Background(), tt.args.contractAddress, tt.args.accountAddress)

			if (err != nil) != tt.wantErr {
				t.Errorf("BaseBalanceOf() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if got.Cmp(&tt.balanceGte) < 0 {
					t.Fatalf("Token = %v, want %d", got, tt.balanceGte.Int64())
				}
			}
		})
	}
}
