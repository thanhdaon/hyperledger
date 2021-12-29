package main

import (
	"fabric-demo/ledgerapi"
	"fabric-demo/phonecard"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	contract := new(phonecard.Contract)
	contract.TransactionContextHandler = new(ledgerapi.TransactionContext)
	contract.Name = "vn.mobifone.phonecard"
	contract.Info.Version = "0.0.1"

	chaincode, err := contractapi.NewChaincode(contract)
	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode: %s", err.Error()))
	}

	chaincode.Info.Title = "PhonecardChaincode"
	chaincode.Info.Version = "0.0.1"

	if err := chaincode.Start(); err != nil {
		panic(fmt.Sprintf("Error starting chaincode: %s", err.Error()))
	}
}
