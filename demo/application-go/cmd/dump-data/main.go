package main

import (
	"asset-transfer-basic/generate"
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(generate.PhoneNumber())
	}

	for _, pc := range generate.Phonecards(10) {
		fmt.Println(pc)
	}
}
