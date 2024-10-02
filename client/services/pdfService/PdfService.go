package pdfService

import (
	"fmt"
	"os"
)

func WriteInDefaultTxt(text string) error {

	file, err := os.Create("DollarQuote.txt")
	if err != nil {
		return err
	}

	_, err = file.WriteString(text)
	if err != nil {
		return err
	}

	fmt.Println("   WRITE TXT VALUE -> ", text)

	return nil
}
