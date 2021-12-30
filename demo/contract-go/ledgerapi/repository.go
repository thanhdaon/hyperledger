package ledgerapi

import (
	"fabric-demo/errors"
	"fabric-demo/phonecard"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Repsitory struct {
	ctx  contractapi.TransactionContextInterface
	name string
}

func NewRepository(ctx contractapi.TransactionContextInterface) Repsitory {
	return Repsitory{ctx: ctx, name: "vn.mobifone.phonecardlist"}
}

func (r Repsitory) SaveCard(pc phonecard.Phonecard) error {
	const op errors.Op = "ledgerapi.Repository.AddCard"

	key, err := r.buildKey(pc.Code())
	if err != nil {
		return errors.E(op, err)
	}

	data, err := toBytes(pc)
	if err != nil {
		return errors.E(op, err)
	}

	if err := r.ctx.GetStub().PutState(key, data); err != nil {
		return errors.E(op, err)
	}

	return nil
}

func (r Repsitory) FindCardByCode(code string) (phonecard.Phonecard, error) {
	const op errors.Op = "ledgerapi.Repository.AddCard"

	key, err := r.buildKey(code)
	if err != nil {
		return phonecard.Nil, errors.E(op, errors.KNotFound, err)
	}

	data, err := r.ctx.GetStub().GetState(key)
	if err != nil {
		return phonecard.Nil, errors.E(op, err)
	}

	pc, err := fromBytes(data)
	if err != nil {
		return phonecard.Nil, errors.E(op, err)
	}

	return pc, nil
}

func (r Repsitory) buildKey(code string) (string, error) {
	const op errors.Op = "ledgerapi.Repository.buildKey"

	key, err := r.ctx.GetStub().CreateCompositeKey(r.name, []string{code})
	if err != nil {
		return "", errors.E(op, err)
	}

	return key, nil
}
