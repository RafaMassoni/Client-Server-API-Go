package responseModel

import (
	"encoding/json"
	"fmt"
)

type DollarQuoteResponse struct {
	DollarValue string `json:"bid"`
}

func ConvertJsonToDollarQuoteResponse(jsonPuro []byte) DollarQuoteResponse {

	var response DollarQuoteResponse

	err := json.Unmarshal(jsonPuro, &response)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON: ", err)
	}

	return response
}
