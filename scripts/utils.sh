#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a collection of bash functions used by different scripts

ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
PEER0_ORG1_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
PEER0_ORG2_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
PEER0_ORG3_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
PEER0_ORG4_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/ca.crt
PEER0_ORG5_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/tls/ca.crt
PEER0_ORG6_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org6.example.com/peers/peer0.org6.example.com/tls/ca.crt
PEER0_ORG7_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org7.example.com/peers/peer0.org7.example.com/tls/ca.crt

CC_VERSION=4.040

# verify the result of the end-to-end test
verifyResult() {
  if [ $1 -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to execute End-2-End Scenario ==========="
    echo
    exit 1
  fi
}


setGlobals() {
  PEER=$1
  ORG=$2
  if [ $ORG -eq 1 ]; then
    CORE_PEER_LOCALMSPID="Org1MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    if [ $PEER -eq 0 ]; then
      CORE_PEER_ADDRESS=peer0.org1.example.com:7051
    else
      CORE_PEER_ADDRESS=peer1.org1.example.com:9051
    fi
  
  elif [ $ORG -eq 2 ]; then
    CORE_PEER_LOCALMSPID="Org2MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
    if [ $PEER -eq 0 ]; then
      CORE_PEER_ADDRESS=peer0.org2.example.com:9051
    else
      CORE_PEER_ADDRESS=peer1.org2.example.com:10051
    fi
  
  elif [ $ORG -eq 3 ]; then
    CORE_PEER_LOCALMSPID="Org3MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
    if [ $PEER -eq 0 ]; then
      CORE_PEER_ADDRESS=peer0.org3.example.com:10051
    else
      CORE_PEER_ADDRESS=peer1.org3.example.com:12051
    fi

  elif [ $ORG -eq 4 ]; then
    CORE_PEER_LOCALMSPID="Org4MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG4_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org4.example.com/users/Admin@org4.example.com/msp
    if [ $PEER -eq 0 ]; then
      CORE_PEER_ADDRESS=peer0.org4.example.com:12051
    else
      CORE_PEER_ADDRESS=peer1.org4.example.com:14051
    fi
  
  elif [ $ORG -eq 5 ]; then
    CORE_PEER_LOCALMSPID="Org5MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG5_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org5.example.com/users/Admin@org5.example.com/msp
    if [ $PEER -eq 0 ]; then
      CORE_PEER_ADDRESS=peer0.org5.example.com:14051
    else
      CORE_PEER_ADDRESS=peer1.org5.example.com:16051
    fi
  
  elif [ $ORG -eq 6 ]; then
    CORE_PEER_LOCALMSPID="Org6MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG6_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org6.example.com/users/Admin@org6.example.com/msp
    if [ $PEER -eq 0 ]; then
      CORE_PEER_ADDRESS=peer0.org6.example.com:16051
    else
      CORE_PEER_ADDRESS=peer1.org6.example.com:18051
    fi
  
  elif [ $ORG -eq 7 ]; then
    CORE_PEER_LOCALMSPID="Org7MSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG7_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org7.example.com/users/Admin@org7.example.com/msp
    if [ $PEER -eq 0 ]; then
      CORE_PEER_ADDRESS=peer0.org7.example.com:18051
    else
      CORE_PEER_ADDRESS=peer1.org7.example.com:20051
    fi
  else
    echo "================== ERROR !!! ORG Unknown =================="
    fi

  if [ "$VERBOSE" == "true" ]; then
    env | grep CORE
  fi
}


updateAnchorPeers() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  set -x
  peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA >&log.txt
  res=$?
  set +x

  cat log.txt
  verifyResult $res "Anchor peer update failed"
  echo "===================== Anchor peers updated for org '$CORE_PEER_LOCALMSPID' on channel '$CHANNEL_NAME' ===================== "
  sleep $DELAY
  echo
}

## Sometimes Join takes time hence RETRY at least 5 times
joinChannelWithRetry() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  echo "The peer: '$PEER' and the org: '$ORG'"

  set -x
  peer channel join -b $CHANNEL_NAME.block >&log.txt
  res=$?
  set +x
  cat log.txt


  if [ $res -ne 0 -a $COUNTER -lt $MAX_RETRY ]; then
    COUNTER=$(expr $COUNTER + 1)
    echo "peer${PEER}.org${ORG} failed to join the channel, Retry after $DELAY seconds"
    sleep $DELAY
    joinChannelWithRetry $PEER $ORG
  else
    COUNTER=1
  fi
  verifyResult $res "After $MAX_RETRY attempts, peer${PEER}.org${ORG} has failed to join channel '$CHANNEL_NAME' "
}

installChaincode() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG
  VERSION=${3:-1.0}
  set -x
  peer chaincode install -n taskmatching -v ${CC_VERSION} -l ${LANGUAGE} -p ${CC_SRC_PATH} >&log.txt
  res=$?
  set +x
  cat log.txt
  verifyResult $res "Chaincode installation on peer${PEER}.org${ORG} has failed"
  echo "===================== Chaincode is installed on peer${PEER}.org${ORG} ===================== "
  echo
}

instantiateChaincode() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG
  VERSION=${3:-1.0}

  # while 'peer chaincode' command can get the orderer endpoint from the peer
  # (if join was successful), let's supply it directly as we know it using
  # the "-o" option
  
  set -x
  peer chaincode instantiate -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -l ${LANGUAGE} -v ${CC_VERSION} -c '{"Args":["init"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer','Org4MSP.peer','Org5MSP.peer','Org6MSP.peer','Org7MSP.peer')"
  res=$?
  set +x

  verifyResult $res "Chaincode instantiation on peer${PEER}.org${ORG} on channel '$CHANNEL_NAME' failed"
  echo "===================== Chaincode is instantiated on peer${PEER}.org${ORG} on channel '$CHANNEL_NAME' ===================== "
  echo
}


# chaincodeInvoke 
# Goes through an example of how to calculate taskmatchings.
chaincodeInvoke() {

  ## Initializing the network with this line here:
  set -x
  echo "Initializing the network:"
  peer chaincode invoke -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["Initialize"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to Initialize network! ==========="
    echo
    exit 1
  fi

  ## Creating a taskmatching to be solved on the network:
  set -x
  echo "Creating a taskmatching:"
  peer chaincode invoke -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["createTaskMatching", "work", "[[1.3,5.4,9.6],[2.8,6.3,10.2],[11.5,7.6,3.7],[12.1,8.4,4.2]]"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to create taskmatching! ==========="
    echo
    exit 1
  fi

  echo "Sleeping for 3 seconds while ledger updating"
  sleep 3

  ## Reading the created taskmatching::
  set -x
  echo "Read the taskmatching"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "work"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi

  ## Read the current peer statuses::
  set -x
  echo "Read the current peer statuses"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "p1"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi

  set -x
  echo "Read the current peer statuses"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "p2"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi

  set -x
  echo "Read the current peer statuses"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "p3"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi

  set -x
  echo "Read the current peer statuses"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "p4"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi

  set -x
  echo "Read the current peer statuses"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "p5"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi

  set -x
  echo "Read the current peer statuses"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "p6"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi


  ##The following setGlobals aren't mandatory however, they represent the different peers running the calculations.
  setGlobals 0 1
  ## Calculate task matchings::
  set -x
  echo "Calculating Task Matching"
  peer chaincode invoke -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["calculateTaskMatching", "p1"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to calculate taskmatching! ==========="
    echo
    exit 1
  fi

  sleep 3

  setGlobals 0 2
  set -x
  echo "Calculating Task Matching"
  peer chaincode invoke -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["calculateTaskMatching", "p2"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to calculate taskmatching! ==========="
    echo
    exit 1
  fi

  sleep 3

  setGlobals 0 3
  set -x
  echo "Calculating Task Matching"
  peer chaincode invoke -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["calculateTaskMatching", "p3"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to calculate taskmatching! ==========="
    echo
    exit 1
  fi

  sleep 3

  setGlobals 0 4
  set -x
  echo "Calculating Task Matching"
  peer chaincode invoke -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["calculateTaskMatching", "p4"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to calculate taskmatching! ==========="
    echo
    exit 1
  fi

  sleep 3

  setGlobals 0 5
  set -x
  echo "Calculating Task Matching"
  peer chaincode invoke -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["calculateTaskMatching", "p5"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to calculate taskmatching! ==========="
    echo
    exit 1
  fi

  sleep 3

  setGlobals 0 6
  set -x
  echo "Calculating Task Matching"
  peer chaincode invoke -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["calculateTaskMatching", "p6"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to calculate taskmatching! ==========="
    echo
    exit 1
  fi

  echo "Sleeping for 3s to make sure ledger updates in time:"
  sleep 3

  ## Read the current peer statuses now that the work has been completed:
  set -x
  echo "Read the current peer statuses"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "p1"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi

  set -x
  echo "Read the current peer statuses"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "p2"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi

  set -x
  echo "Read the current peer statuses"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "p3"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi

  set -x
  echo "Read the current peer statuses"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "p4"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi

    set -x
  echo "Read the current peer statuses"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "p5"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi

    set -x
  echo "Read the current peer statuses"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "p6"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi
  #Must call calculateTaskMatching method again in order for the ledger to realize that all tasks are complete (Doesn't matter which peer it's called on p1/p2/p3 all work) :
  set -x
  echo "Picking optimal taskmatching and writing the solution to the ledger"
  peer chaincode invoke -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["calculateTaskMatching", "p3"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to calculate the optimal taskmatching! ==========="
    echo
    exit 1
  fi

  echo "Sleeping for 3s while ledger updates"
  sleep 3


  setGlobals 0 1
  #Can now view the solution that has been created for the given taskmatching:
  set -x
  echo "viewing solution"
  peer chaincode query -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n taskmatching -c '{"Args":["readTaskMatching", "1"]}'
  res=$?
  set +x

  if [ $res -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to read taskmatching! ==========="
    echo
    exit 1
  fi

  echo "===================== Invoke transaction successful on $PEERS on channel '$CHANNEL_NAME' ===================== "
  echo
}
