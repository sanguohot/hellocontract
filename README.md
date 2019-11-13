# hellocontract
simple solidty hello contract
## 环境依赖
go1.12+
## 克隆源码
```
# git clone https://github.com/sanguohot/hellocontract
# cd hellocontract
# go build --ldflags "-linkmode external -extldflags -static" -o hellocontract.exe main.go
# go build --ldflags "-linkmode external -extldflags -static" -o hellocontract main.go
```
## 编译（linux）
```
# go build --ldflags "-linkmode external -extldflags -static" -o hellocontract main.go
```
## 编译（windows）
```
# go build --ldflags "-linkmode external -extldflags -static" -o hellocontract.exe main.go
```
## 运行（windows）
```
# hellocontract.exe -h
  Usage of hellocontract.exe:
    -action string
          blockchain hello contract action, get or put(update) (default "get")
    -contract-addr string
          blockchain hello contract address (default "0xfcd14ED03E6D94CA127d557a1883Dd042a81ea11")
    -data string
          update hello contract with this data
    -private-key string
          using private key to commit a transaction, it will change data.
    -rpc-addr string
          blockchain rpc address (default "10.0.252.5:8546")
# hellocontract.exe -action get
  2019-11-13T15:24:39.429+0800    info    hello/hello.go:38       query hello contract get >>>>>>>>> 4444
# hellocontract.exe -action put -data "hello, my name is sgh" -private-key bcec428d5205abe0f0cc8a734083908d9eb8563e31f943d760786edf42ad67dd
  2019-11-13T15:35:03.494+0800    info    hello/hello.go:80       update hello contract transaction hash >>>>>>>>> 0x800887886a91dae3a537e61df2e60aa1d01bb9b56b5353f2abd5efd6c4bbaea1
# hellocontract.exe -action get
  2019-11-13T15:35:24.500+0800    info    hello/hello.go:38       query hello contract get >>>>>>>>> hello, my name is sgh
```