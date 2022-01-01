package contract

import (
	"fabric-demo/phonecard"
)

type Phonecard struct {
	Code                 string `json:"code"`
	Activated            bool   `json:"activated"`
	ActivatedPhoneNumber string `json:"activatedPhoneNumber,omitempty" metadata:",optional"`
	FaceValue            int    `json:"facevalue"`
	IssuedAt             string `json:"issuedAt"`
	ActivatedAt          string `json:"activatedAt,omitempty" metadata:",optional"`
	ExpiredAt            string `json:"expiredAt"`
}

func toPhonecard(entity phonecard.Phonecard) Phonecard {
	ret := Phonecard{
		Code:                 entity.Code(),
		Activated:            entity.Activated(),
		ActivatedPhoneNumber: entity.ActivatedPhoneNumber(),
		FaceValue:            entity.FaceValue(),
		IssuedAt:             entity.IssuedAt().Format("Mon 02, 2006 15:04:05 PM"),
	}

	if isZero := entity.IssuedAt().IsZero(); !isZero {
		ret.IssuedAt = entity.IssuedAt().Format("Mon 02, 2006 15:04:05 PM")
	}

	if isZero := entity.ActivatedAt().IsZero(); !isZero {
		ret.ActivatedAt = entity.ActivatedAt().Format("Mon 02, 2006 15:04:05 PM")
	}

	if isZero := entity.ExpiredAt().IsZero(); !isZero {
		ret.ExpiredAt = entity.ExpiredAt().Format("Mon 02, 2006 15:04:05 PM")
	}

	return ret
}

func toPhonecards(entities []phonecard.Phonecard) []Phonecard {
	ret := []Phonecard{}

	for _, e := range entities {
		ret = append(ret, toPhonecard(e))
	}

	return ret
}
