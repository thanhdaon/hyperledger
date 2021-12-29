package phonecard

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct {
	contractapi.Contract
}

func (c *Contract) Instantiate() {
	fmt.Println("Instantiated")
}

func (c *Contract) Issue(ctx TransactionContext, code string, facevalue int, duration string) error {
	return nil
}

func (c *Contract) ActiveCard(ctx TransactionContext, code string, phonenumber string) error {
	return nil
}
