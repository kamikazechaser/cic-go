package cic_net

import (
	"os"
)

type tConfig struct {
	rpcProvider string
	tokenIndex  string
}

var conf = &tConfig{
	rpcProvider: os.Getenv("RPC_PROVIDER"),
	tokenIndex:  os.Getenv("TOKEN_INDEX"),
}
