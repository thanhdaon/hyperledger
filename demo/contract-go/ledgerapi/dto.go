package ledgerapi

import (
	"encoding/json"
	"fabric-demo/errors"
	"fabric-demo/phonecard"
	"time"
)

type Phonecard struct {
	Code                 string    `json:"code"`
	Activated            bool      `json:"activated"`
	ActivatedPhoneNumber string    `json:"activatedPhoneNumber"`
	FaceValue            int       `json:"facevalue"`
	IssuedAt             time.Time `json:"issuedAt"`
	ActivatedAt          time.Time `json:"activatedAt"`
	ExpiredAt            time.Time `json:"expiredAt"`
}

func toBytes(pc phonecard.Phonecard) ([]byte, error) {
	const op errors.Op = "ledgerapi.dto.toBytes"

	dto := Phonecard{
		Code:                 pc.Code(),
		Activated:            pc.Activated(),
		ActivatedPhoneNumber: pc.ActivatedPhoneNumber(),
		FaceValue:            pc.FaceValue(),
		IssuedAt:             pc.IssuedAt(),
		ActivatedAt:          pc.ActivatedAt(),
		ExpiredAt:            pc.ExpiredAt(),
	}

	ret, err := json.Marshal(dto)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return ret, nil
}

func fromBytes(data []byte) (phonecard.Phonecard, error) {
	const op errors.Op = "ledgerapi.dto.fromBytes"

	var dto Phonecard

	if err := json.Unmarshal(data, &dto); err != nil {
		return phonecard.Nil, errors.E(op, err)
	}

	pc, err := phonecard.FromPersistenceLayer(
		dto.Code,
		dto.Activated,
		dto.ActivatedPhoneNumber,
		dto.FaceValue,
		dto.IssuedAt,
		dto.ActivatedAt,
		dto.ExpiredAt,
	)

	if err != nil {
		return phonecard.Nil, errors.E(op, err)
	}

	return pc, nil
}
