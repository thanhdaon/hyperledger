package contract

import (
	ledgerapi "fabric-demo/ledger-api"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TransactionContext interface {
	contractapi.TransactionContextInterface
	Repository() ledgerapi.Repsitory
}
