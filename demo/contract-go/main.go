package main

import (
	"fabric-demo/contract"
	"fabric-demo/ledgerapi"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	ct := new(contract.Contract)
	ct.TransactionContextHandler = new(ledgerapi.TransactionContext)
	ct.Name = "vn.mobifone.phonecard"
	ct.Info.Version = "0.0.1"

	chaincode, err := contractapi.NewChaincode(ct)
	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode: %s", err.Error()))
	}

	chaincode.Info.Title = "PhonecardChaincode"
	chaincode.Info.Version = "0.0.1"

	if err := chaincode.Start(); err != nil {
		panic(fmt.Sprintf("Error starting chaincode: %s", err.Error()))
	}
}
