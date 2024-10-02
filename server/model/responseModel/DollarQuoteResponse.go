package responseModel

import (
	"encoding/json"
	"log"
)

type EconomiaResponse struct {
	USDBRL struct {
		DollarValue string `json:"bid"`
	} `json:"USDBRL"`
}

type DollarQuoteResponse struct {
	DollarValue string `json:"bid"`
}

func ConvertJsonToDollarQuoteResponse(jsonPuro []byte) EconomiaResponse {

	var response EconomiaResponse

	err := json.Unmarshal(jsonPuro, &response)
	if err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	return response
}

func ConvertEconomiaResponseToDollarQuoteResponse(economiaResponse EconomiaResponse) DollarQuoteResponse {

	response := DollarQuoteResponse{
		DollarValue: economiaResponse.USDBRL.DollarValue,
	}

	return response
}
