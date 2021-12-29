package phonecard

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type TransactionContext interface {
	contractapi.TransactionContextInterface
	GetRepository() Repository
}
