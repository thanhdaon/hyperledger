package ledgerapi

import (
	"fabric-demo/phonecard"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TransactionContext struct {
	contractapi.TransactionContext
	repository phonecard.Repository
}

func (tc *TransactionContext) GetRepository() phonecard.Repository {
	if tc.repository == nil {
		tc.repository = newRepository(tc)
	}

	return tc.repository
}
