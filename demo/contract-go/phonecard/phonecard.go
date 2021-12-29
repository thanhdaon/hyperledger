package phonecard

import "time"

type Phonecard struct {
	code                 string
	activated            bool
	activatedPhoneNumber string
	facevalue            int
	issuedAt             time.Time
	activatedAt          time.Time
	expiredAt            time.Time
}

func (pc *Phonecard) Active(phoneNumber string) {
	pc.activated = true
	pc.activatedPhoneNumber = phoneNumber
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
