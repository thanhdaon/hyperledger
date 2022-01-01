package main

import (
	"asset-transfer-basic/smartcontract"
	"fmt"
	"log"
)

func main() {
	smartcontract, err := smartcontract.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = smartcontract.SubmitTransaction("ActiveCard", "102144116655", "0128374483")
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(" [INFO] active card done! \n")
}
