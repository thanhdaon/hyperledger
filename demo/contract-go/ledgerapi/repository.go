package ledgerapi

import (
	"fabric-demo/phonecard"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Repsitory struct {
	ctx  contractapi.TransactionContextInterface
	name string
}

func newRepository(ctx contractapi.TransactionContextInterface) Repsitory {
	return Repsitory{ctx: ctx, name: "vn.mobifone.phonecardlist"}
}

func (r Repsitory) AddCard(pc phonecard.Phonecard) error {
	attributes := []string{pc.Code()}

	key, err := r.ctx.GetStub().CreateCompositeKey(r.name, attributes)
	if err != nil {
		return err
	}

	data, err := toBytes(pc)
	if err != nil {
		return err
	}

	return r.ctx.GetStub().PutState(key, data)
}

func (r Repsitory) ActiveCard(code string, phoneNumber string) error {
	return nil
}
