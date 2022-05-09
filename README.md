[![Go Reference](https://pkg.go.dev/badge/github.com/grassrootseconomics/cic_go.svg)](https://pkg.go.dev/github.com/grassrootseconomics/cic_go)
[![Go](https://github.com/grassrootseconomics/cic_go/actions/workflows/go.yml/badge.svg)](https://github.com/grassrootseconomics/cic_go/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/grassrootseconomics/cic_go/badge.svg?branch=sohail/cic_net_updates)](https://coveralls.io/github/grassrootseconomics/cic_go?branch=sohail/cic_net_updates)

# cic-go
Go modules to access various parts of the cic stack

## Implemented

- meta (web2 metadata store)
- net (cic smart contracts)

## Installation

`go get -u github.com/grassrootseconomics/cic-go`

## Usage

```go
import (
    cic_meta "github.com/grassrootseconomics/cic-go/meta"
    cic_net  "github.com/grassrootseconomics/cic-go/net"
) 
```