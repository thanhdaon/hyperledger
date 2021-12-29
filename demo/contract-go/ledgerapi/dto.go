package ledgerapi

import (
	"encoding/json"
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
		return nil, err
	}

	return ret, nil
}
