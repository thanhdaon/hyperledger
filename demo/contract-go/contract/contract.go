package contract

import (
	"fabric-demo/errors"
	"fabric-demo/phonecard"
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
	const op errors.Op = "Contract.Issue"

	pc, err := phonecard.New(code, facevalue, duration)
	if err != nil {
		return errors.E(op, err)
	}

	if err := ctx.Repository().SaveCard(pc); err != nil {
		return errors.E(op, err)
	}

	return nil
}

func (c *Contract) ActiveCard(ctx TransactionContext, code string, phoneNumber string) error {
	const op errors.Op = "Contract.ActiveCard"

	pc, err := ctx.Repository().FindCardByCode(code)
	if err != nil {
		return errors.E(op, err)
	}

	pc.Active(phoneNumber)

	if err := ctx.Repository().SaveCard(pc); err != nil {
		return errors.E(op, err)
	}

	return nil
}
