# HYPERLEDGER FABRIC APPLICATION

## Generating Certs

STEP 1. Start CA Server

```shell
docker-compose -f docker-compose-ca.yaml up
```

STEP 2. Genereate Certification.

```shell
docker exec -it ca-client /etc/hyperledger/fabric-ca-client/scripts/gen-certs.sh
```

STEP 3.

```shell
docker cp ca-client:/etc/hyperledger/fabric-ca-client/crypto-config ./crypto-config
```


## Generating Network Artifacts

STEP 1. Generating Artifacts

```shell

configtxgen -profile OrdererGenesis -channelID syschannel -outputBlock ./orderer/genesis.block
configtxgen -profile MainChannel -outputCreateChannelTx ./channels/mainchannel.tx -channelID mainchannel
configtxgen -profile MainChannel -outputAnchorPeersUpdate ./channels/org1-anchors.tx -channelID mainchannel -asOrg org1
configtxgen -profile MainChannel -outputAnchorPeersUpdate ./channels/org2-anchors.tx -channelID mainchannel -asOrg org2

```

STEP 2. Start services.

```shell
docker-compose up
```

STEP 3. Create Channel

```shell

docker exec cli-peer0-org1 bash -c 'peer channel create -c mainchannel -f ./channels/mainchannel.tx -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem'

```

STEP 4. Join block

```shell

docker exec cli-peer0-org1 bash -c 'peer channel join -b mainchannel.block'
docker exec cli-peer0-org1 bash -c 'cp mainchannel.block ./channels/mainchannel.block'
docker exec cli-peer1-org1 bash -c 'peer channel join -b ./channels/mainchannel.block'
docker exec cli-peer0-org2 bash -c 'peer channel join -b ./channels/mainchannel.block'
docker exec cli-peer1-org2 bash -c 'peer channel join -b ./channels/mainchannel.block'

```

STEP 4. Update the transection to encrypt

```shell

docker exec cli-peer0-org1 bash -c 'peer channel update -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem -c mainchannel -f ./channels/org1-anchors.tx'
docker exec cli-peer0-org2 bash -c 'peer channel update -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem -c mainchannel -f ./channels/org2-anchors.tx'

```

STEP 5. Install Chaincode
***please wait step 4 15 second***

```shell

docker exec cli-peer0-org1 bash -c 'peer chaincode install -p rawresources -n rawresources -v 0'
docker exec cli-peer1-org1 bash -c 'peer chaincode install -p rawresources -n rawresources -v 0'
docker exec cli-peer0-org2 bash -c 'peer chaincode install -p rawresources -n rawresources -v 0'
docker exec cli-peer1-org2 bash -c 'peer chaincode install -p rawresources -n rawresources -v 0'

```

STEP 6. Test Network

```shell
docker exec cli-peer0-org1 bash -c "peer chaincode instantiate -C mainchannel -n rawresources -v 0 -c '{\"Args\":[]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem"
```

***in contrainner***

Store Data

```shell
docker exec -it cli-peer0-org1 bash

peer chaincode invoke -C mainchannel -n rawresources -c '{"Args":["store", "{\"id\": \"1\",\"name\":\"Iron Ore\",\"weight\":42000}"]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem

peer chaincode invoke -C mainchannel -n rawresources -c '{"Args":["store", "{\"id\": \"2\",\"name\":\"Iron Ore\",\"weight\":20000}"]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem

peer chaincode invoke -C mainchannel -n rawresources -c '{"Args":["store", "{\"id\": \"3\",\"name\":\"Iron Ore\",\"weight\":10000}"]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem

```

Query Data

```shell
docker exec cli-peer0-org1 bash -c "peer chaincode query -C mainchannel -n rawresources -c '{\"Args\":[\"index\",\"\",\"\"]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem"

docker exec cli-peer1-org1 bash -c "peer chaincode query -C mainchannel -n rawresources -c '{\"Args\":[\"index\",\"\",\"\"]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem"

docker exec cli-peer0-org2 bash -c "peer chaincode query -C mainchannel -n rawresources -c '{\"Args\":[\"index\",\"\",\"\"]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem"

docker exec cli-peer1-org2 bash -c "peer chaincode query -C mainchannel -n rawresources -c '{\"Args\":[\"index\",\"\",\"\"]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem"
```

Update Data

```shell
docker exec -it cli-peer0-org1 bash

peer chaincode invoke -C mainchannel -n rawresources -c '{"Args":["update", "3",  "{\"id\": \"3\",\"name\":\"Iron Ore\",\"weight\":20000, \"visible\": true }"]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem
```

## Update Chaincode

GET Current Block

```shell

check from channel

docker exec cli-peer0-org1 bash -c "peer channel getinfo -c mainchannel --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem"

or check from log

docker logs peer0-org1-service 2>&1 |grep 'Received block'
```

```shell
docker exec cli-peer0-org1 bash -c 'peer chaincode install -p rawresources -n rawresources -v 2'
docker exec cli-peer1-org1 bash -c 'peer chaincode install -p rawresources -n rawresources -v 2'
docker exec cli-peer0-org2 bash -c 'peer chaincode install -p rawresources -n rawresources -v 2'
docker exec cli-peer1-org2 bash -c 'peer chaincode install -p rawresources -n rawresources -v 2'
```

*** Upgrade Channel

```shell
docker exec cli-peer0-org1 bash -c "peer chaincode upgrade -C mainchannel -n rawresources -v 2 -c '{"Args":[]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem"
```

* Remark testing invoke or query

```shell

docker exec -it cli-peer0-org1 bash

peer chaincode invoke -C mainchannel -n rawresources -c '{"Args":["store", "{"id":2,"name":"Iron Ore","weight":20000}"]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem peer chaincode invoke -C mainchannel -n rawresources -c '{"Args":["store", "{"id":3,"name":"Iron Ore","weight":10000}"]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem

peer chaincode query -C mainchannel -n rawresources -c '{"Args":["queryString", "{"selector":{ "weight": { "$gt":5000 } }}"]}' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/ca-root-7054.pem

```
