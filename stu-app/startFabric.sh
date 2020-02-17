#!/bin/bash
set -e

starttime=$(date +%s)
# clean the keystore
if [ ! -d ~/.hfc-key-store/ ]; then
	mkdir ~/.hfc-key-store/
fi

# launch network; create channel and join peer to channel
cd ../basic-network
./start.sh

# launch the CLI container ,install chaincode
docker-compose -f ./docker-compose.yml up -d cli

docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode install -n stu-app -v 1.0 -p github.com/stu-app
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n stu-app -v 1.0 -c '{"Args":[""]}' -P "OR ('Org1MSP.member','Org2MSP.member')"
sleep 3
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n stu-app -c '{"function":"initLedger","Args":[""]}'

# enroll admin and users
cd ../stu-app
node registerAdmin.js
printf "\n-------registerAdmin success--------\n\n"
node registerUser.js
printf "\n-------registerUser success-------\n\n"

echo ==============================================================================
printf "\n-------Total execution time : $(($(date +%s) - starttime)) secs ...\n\n"
printf "\n-------Start blockchain network success-------\n\n"
echo =====================================================
echo ======= Students Grades Fabric System Started =======
echo =====================================================
echo please run "node server.js" and visit localhost:8000
echo =====================================================
