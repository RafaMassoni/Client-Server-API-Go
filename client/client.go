package main

import (
	"client/clients/serverClient"
	"client/services/pdfService"
	"fmt"
)

func main() {
	fmt.Println("\n ########## Client iniciando ########## \n ")
	defer fmt.Println("\n\n ########## Client finalizado ########## \n ")

	dollarQuote, err := serverClient.GetDollarQuote()
	if err != nil {
		fmt.Print("Oops, ocorreu um erro: ", err)
		return
	}

	fmt.Println("REQUEST SUCCESSFUL")

	err = pdfService.WriteInDefaultTxt("Dollar " + dollarQuote.DollarValue)
	if err != nil {
		fmt.Println(err)
		return
	}

}
