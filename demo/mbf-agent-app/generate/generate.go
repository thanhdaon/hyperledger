package generate

import "github.com/Pallinder/go-randomdata"

type Phonecard struct {
	Code      string `json:"code"`
	FaceValue string `json:"facevalue"`
	Duration  string `json:"duration"`
}

func Phonecards(n int) []Phonecard {
	phonecards := []Phonecard{}

	for i := 0; i < n; i++ {
		phonecards = append(phonecards, Phonecard{
			Code:      randomdata.StringNumber(6, ""),
			FaceValue: randomdata.StringSample("10000", "20000", "50000", "100000", "200000"),
			Duration:  "1h",
		})
	}

	return phonecards
}

func PhoneNumber() string {
	return "84" + randomdata.StringNumber(5, "")
}
