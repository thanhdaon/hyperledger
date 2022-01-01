package main

import (
	"asset-transfer-basic/generate"
	"asset-transfer-basic/smartcontract"
	"fmt"
	"log"
)

func main() {
	smartcontract, err := smartcontract.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	for _, pc := range generate.Phonecards(6) {
		fmt.Println(pc)
		_, err := smartcontract.SubmitTransaction("Issue", pc.Code, pc.FaceValue, pc.Duration)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf(" [INFO] success dump card %v \n", pc)
	}

	fmt.Printf(" [INFO] Dump data done! \n")
}
