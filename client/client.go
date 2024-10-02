package main

import (
	"client/clients/serverClient"
	"client/services/pdfService"
	"fmt"
)

func main() {
	fmt.Println("---- CLIENT START ----")
	defer fmt.Println("\n---- CLIENT END ----")

	dollarQuote, err := serverClient.GetDollarQuote()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("REQUEST SUCCESSFUL")

	err = pdfService.WriteInDefaultTxt("Dollar " + dollarQuote.DollarValue)
	if err != nil {
		fmt.Println(err)
		return
	}

}
