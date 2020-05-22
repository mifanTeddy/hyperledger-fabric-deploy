# 部署文档

## 设置工作路径

```bash
export FABRIC_CFG_PATH=$GOPATH/src/github.com/hyperledger/fabric/haierbiomedical/deploy
```

## 环境清理

```bash
rm -fr config/*
rm -fr crypto-config/*
```

## 生成证书文件

```bash
cryptogen generate --config=./crypto-config.yaml
```

## 生成创世区块

```bash
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./config/genesis.block
```

## 生成通道的创世交易

```bash
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./config/examplechannel.tx -channelID examplechannel
```

## 生成组织关于通道的锚节点（主节点）交易

```bash
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/Org1MSPanchors.tx -channelID examplechannel -asOrg Org1MSP
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/Org2MSPanchors.tx -channelID examplechannel -asOrg Org2MSP
```

## 创建通道

```bash
peer channel create -o orderer.example.com:7050 -c examplechannel -f /etc/hyperledger/config/examplechannel.tx
```

## 加入通道

```bash
peer channel join -b examplechannel.block
```

## 设置主节点

```bash
peer channel update -o orderer.example.com:7050 -c examplechannel -f /etc/hyperledger/config/Org1MSPanchors.tx
```

## 链码安装

```bash
peer chaincode install -n examplechaincode -v 1.0.0 -l golang -p github.com/chaincode/examplechaincode
```

## 链码实例化

```bash
peer chaincode instantiate -o orderer.haierbiomedical.com:7050 -C examplechannel -n examplechaincode -l golang -v 1.0.0 -c '{"Args":["init"]}'
```

## 链码交互

```bash
peer chaincode invoke -C examplechannel -n examplechaincode -c '{"Args":["set", "example1", "user222xxx"]}'

```

## 链码升级

```bash
peer chaincode install -n examplechaincode -v 1.0.2 -l golang -p github.com/chaincode/examplechaincode
peer chaincode upgrade -C examplechannel -n examplechaincode -v 1.0.2 -c '{"Args":[""]}'
```

## 链码查询

```bash
peer chaincode invoke -n examplechaincode -c '{"Args":["set", "vvv", "{67789}"]}' -C examplechannel
peer chaincode invoke -n examplechaincode -c '{"Args":["set", "example1", "1233user xxx"]}' -C examplechannel
peer chaincode query -n examplechaincode -c '{"Args":["get", "vvv"]}' -C examplechannel
peer chaincode query -n examplechaincode -c '{"Args":["getHistory", "567876545678"]}' -C examplechannel
peer chaincode query -n examplechaincode  -C examplechannel -c '{"Args":["get", "example1"]}'
```
