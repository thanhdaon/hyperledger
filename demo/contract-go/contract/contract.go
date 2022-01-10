package contract

import (
	"fabric-demo/errors"
	"fabric-demo/phonecard"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct {
	contractapi.Contract
}

func (c *Contract) Instantiate() {
	fmt.Println("Instantiated Hello")
}

func (c *Contract) Issue(ctx TransactionContext, code string, facevalue int, duration string) error {
	const op errors.Op = "Contract.Issue"

	mspID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return errors.E(op, err)
	}

	if mspID != "Org1MSP" {
		return errors.E(op, fmt.Errorf("Permissiondenied"))
	}

	ts, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return errors.E(op, err)
	}

	issuedAt := time.Unix(ts.Seconds, 0)

	pc, err := phonecard.New(code, facevalue, duration, issuedAt)
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

	mspID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return errors.E(op, err)
	}

	if mspID != "Org1MSP" {
		return errors.E(op, fmt.Errorf("Permissiondenied"))
	}

	pc, err := ctx.Repository().FindCardByCode(code)
	if err != nil {
		return errors.E(op, err)
	}

	if pc.Activated() {
		return errors.E(op, errors.KCardActivated, fmt.Errorf("Card already activated"))
	}

	ts, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return errors.E(op, err)
	}

	pc.Active(phoneNumber, time.Unix(ts.Seconds, 0))

	if err := ctx.Repository().SaveCard(pc); err != nil {
		return errors.E(op, err)
	}

	return nil
}

func (c *Contract) GetAllCards(ctx TransactionContext) ([]Phonecard, error) {
	const op errors.Op = "Contract.GetAllCards"

	phonecards, err := ctx.Repository().FindAllCards()
	if err != nil {
		return nil, errors.E(op, err)
	}

	return toPhonecards(phonecards), nil
}

func (c *Contract) PruneAllStates(ctx TransactionContext) error {
	const op errors.Op = "Contract.PruneAllStates"

	if err := ctx.Repository().PruneAllStates(); err != nil {
		return errors.E(op, err)
	}

	return nil
}
