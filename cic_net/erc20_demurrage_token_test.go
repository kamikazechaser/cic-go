package cic_net

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

func TestCicNet_DemurrageToken_ChangePeriod(t *testing.T) {
	type args struct {
		writeTx WriteTx
	}

	// Bootstrap signer
	privateKey, err := crypto.HexToECDSA(conf.privateKey)
	if err != nil {
		t.Fatalf("ECDSA error = %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatalf("ECDSA error = %v", err)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	tests := []struct {
		name       string
		args       args
		wantErr    bool
		balanceGte big.Int
	}{
		{
			name: "ChangePeriod for Sarafu",
			args: args{
				writeTx: WriteTx{
					from:       fromAddress,
					to:         w3.A("0xaB89822F31c2092861F713F6F34bd6877a8C1878"),
					gasLimit:   12000000,
					privateKey: *privateKey,
				},
			},
			wantErr: false,
		},
	}

	for _, testcase := range tests {
		tt := testcase

		t.Run(tt.name, func(t *testing.T) {
			cicnet, err := NewCicNet(conf.rpcProvider, w3.A(conf.tokenIndex))

			if err != nil {
				t.Fatalf("NewCicNet error = %v", err)
			}

			tx, err := cicnet.ChangePeriod(context.Background(), tt.args.writeTx)
			t.Logf("ChangePeriod tx_hash %s", tx.String())

			if (err != nil) != tt.wantErr {
				t.Errorf("ChangePeriod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCicNet_DemurrageToken_ApplyDemurrageLimited(t *testing.T) {
	type args struct {
		rounds  int64
		writeTx WriteTx
	}

	// Bootstrap signer
	privateKey, err := crypto.HexToECDSA(conf.privateKey)
	if err != nil {
		t.Fatalf("ECDSA error = %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatalf("ECDSA error = %v", err)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	tests := []struct {
		name       string
		args       args
		wantErr    bool
		balanceGte big.Int
	}{
		{
			name: "ChangePeriod for Sarafu",
			args: args{
				rounds: 1000,
				writeTx: WriteTx{
					from:       fromAddress,
					to:         w3.A("0xaB89822F31c2092861F713F6F34bd6877a8C1878"),
					gasLimit:   12000000,
					privateKey: *privateKey,
				},
			},
			wantErr: false,
		},
	}

	for _, testcase := range tests {
		tt := testcase

		t.Run(tt.name, func(t *testing.T) {
			cicnet, err := NewCicNet(conf.rpcProvider, w3.A(conf.tokenIndex))

			if err != nil {
				t.Fatalf("NewCicNet error = %v", err)
			}

			tx, err := cicnet.ApplyDemurrageLimited(context.Background(), tt.args.rounds, tt.args.writeTx)
			t.Logf("ApplyDemurrageLimited tx_hash %s", tx.String())

			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyDemurrageLimited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
