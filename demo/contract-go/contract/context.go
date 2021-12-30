package contract

import (
	"fabric-demo/ledgerapi"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TransactionContext interface {
	contractapi.TransactionContextInterface
	Repository() ledgerapi.Repsitory
}
