package ledgerapi

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var nilRepository = Repsitory{}

type TransactionContext struct {
	contractapi.TransactionContext
	repository Repsitory
}

func (tc *TransactionContext) Repository() Repsitory {
	if tc.repository == nilRepository {
		tc.repository = NewRepository(tc)
	}

	return tc.repository
}
