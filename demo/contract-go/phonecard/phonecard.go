package phonecard

import (
	"fabric-demo/errors"
	"fmt"
	"time"
)

var Nil = Phonecard{}

type Phonecard struct {
	code                 string
	activated            bool
	activatedPhoneNumber string
	facevalue            int
	issuedAt             time.Time
	activatedAt          time.Time
	expiredAt            time.Time
}

func New(code string, facevalue int, duration string) (Phonecard, error) {
	const op errors.Op = "phonecard.New"

	if code == "" {
		return Nil, errors.E(op, errors.KBadInput, fmt.Errorf("empty code!"))
	}

	if facevalue == 0 {
		return Nil, errors.E(op, errors.KBadInput, fmt.Errorf("facevalue is zero!"))
	}

	d, err := time.ParseDuration(duration)
	if err != nil {
		return Nil, errors.E(op, errors.KBadInput, err)
	}

	return Phonecard{
		code:                 code,
		activated:            false,
		activatedPhoneNumber: "",
		facevalue:            facevalue,
		issuedAt:             time.Now(),
		expiredAt:            time.Now().Add(d),
	}, nil
}

func FromPersistenceLayer(code string, activated bool, phonenumber string, facevalue int, issuedAt, activatedAt, expiredAt time.Time) (Phonecard, error) {
	const op errors.Op = "phonecard.FromPersistenceLayer"

	if code == "" {
		return Nil, errors.E(op, errors.KBadInput, fmt.Errorf("empty code!"))
	}

	if facevalue == 0 {
		return Nil, errors.E(op, errors.KBadInput, fmt.Errorf("facevalue is zero!"))
	}

	if issuedAt.IsZero() {
		return Nil, errors.E(op, errors.KBadInput, fmt.Errorf("issuedAt is zero!"))
	}

	if expiredAt.IsZero() {
		return Nil, errors.E(op, errors.KBadInput, fmt.Errorf("expiredAt is zero!"))
	}

	if expiredAt.Before(issuedAt) {
		return Nil, errors.E(op, errors.KBadInput, fmt.Errorf("expiredAt is set before issuedAt!"))
	}

	return Phonecard{
		code:                 code,
		activated:            activated,
		activatedPhoneNumber: phonenumber,
		facevalue:            facevalue,
		issuedAt:             issuedAt,
		activatedAt:          activatedAt,
		expiredAt:            expiredAt,
	}, nil
}

func (pc *Phonecard) Active(phoneNumber string) {
	pc.activated = true
	pc.activatedPhoneNumber = phoneNumber
	pc.activatedAt = time.Now()
}

func (pc *Phonecard) Code() string {
	return pc.code
}

func (pc *Phonecard) Activated() bool {
	return pc.activated
}

func (pc *Phonecard) FaceValue() int {
	return pc.facevalue
}

func (pc *Phonecard) IssuedAt() time.Time {
	return pc.issuedAt
}

func (pc *Phonecard) ActivatedAt() time.Time {
	return pc.activatedAt
}

func (pc *Phonecard) ExpiredAt() time.Time {
	return pc.expiredAt
}

func (pc *Phonecard) ActivatedPhoneNumber() string {
	return pc.activatedPhoneNumber
}

func (pc *Phonecard) Expired() bool {
	return time.Now().After(pc.expiredAt)
}
