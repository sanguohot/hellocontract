package hello

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sanguohot/log/v2"
	"github.com/sanguohot/medichain/contracts/hello" // for demo
	"math/big"
	"math/rand"
	"time"
)

func Get(rpcAddr, contactAddress string) {
	if rpcAddr == "" || contactAddress == "" {
		log.Logger.Fatal(errors.New("param rpc-addr/contract-addr required").Error())
	}
	client, err := ethclient.Dial(fmt.Sprintf("http://%s", rpcAddr))
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	if !common.IsHexAddress(contactAddress) {
		log.Logger.Fatal(err.Error())
	}
	address := common.HexToAddress(contactAddress)
	instance, err := hello.NewHello(address, client)
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	result, err := instance.Speak(nil)
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	log.Sugar.Infof("query hello contract get >>>>>>>>> %s", result)
}

func Put(rpcAddr, contactAddress, privateKeyStr, data string) {
	if rpcAddr == "" || contactAddress == "" || privateKeyStr == "" || data == "" {
		log.Logger.Fatal(errors.New("param rpc-addr/contract-addr/private-key/data required").Error())
	}
	client, err := ethclient.Dial(fmt.Sprintf("http://%s", rpcAddr))
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	if !common.IsHexAddress(contactAddress) {
		log.Logger.Fatal(err.Error())
	}
	address := common.HexToAddress(contactAddress)
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Logger.Fatal(err.Error())
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	auth := bind.NewKeyedTransactor(privateKey) // 设置私钥和签名方法
	rand.Seed(time.Now().Unix())
	auth.Nonce = big.NewInt(rand.Int63n(100000000)) // 也是fisco-bcos的randomid字段，这里按照web3.js计算随机数
	auth.Value = nil
	auth.GasLimit = uint64(1000000)
	//auth.BlockLimit = uint64(740) // 区别go-etherneum，似乎是fisco-bcos判断了范围？
	auth.GasPrice = nil
	auth.From = fromAddress
	instance, err := hello.NewHello(address, client)
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	tx, err := instance.SaySomethingElse(auth, data) // 调用Hello.sol 的 SaySomethingElse方法
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	time.Sleep(time.Millisecond * 1100)
	log.Sugar.Infof("update hello contract transaction hash >>>>>>>>> %s", tx.Hash().Hex())
}
