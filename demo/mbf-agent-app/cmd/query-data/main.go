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

	response, err := smartcontract.EvaluateTransaction("GetAllCards")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(" [INFO] %s \n", string(response))
}
