package main

import (
	"flag"
	"github.com/sanguohot/hellocontract/pkg/hello"
	"github.com/sanguohot/log/v2"
)

const (
	actionTypeGet = "get"
	actionTypePut = "put"
)

var (
	bcRpcAddress           string
	bcContractHelloAddress string
	privateKey             string
	action                 string
	data                   string
)

func init() {
	flag.StringVar(&action, "action", "get", "blockchain hello contract action, get or put(update)")
	flag.StringVar(&bcRpcAddress, "rpc-addr", "10.0.252.5:8546", "blockchain rpc address")
	flag.StringVar(&bcContractHelloAddress, "contract-addr", "0xfcd14ED03E6D94CA127d557a1883Dd042a81ea11", "blockchain hello contract address")
	flag.StringVar(&privateKey, "private-key", "", "using private key to commit a transaction, it will change data.")
	flag.StringVar(&data, "data", "", "update hello contract with this data")
	flag.Parse()
}

func main() {
	switch action {
	case actionTypeGet:
		hello.Get(bcRpcAddress, bcContractHelloAddress)
	case actionTypePut:
		hello.Put(bcRpcAddress, bcContractHelloAddress, privateKey, data)
	default:
		log.Sugar.Warnf("unknown action %s", action)
	}
}
