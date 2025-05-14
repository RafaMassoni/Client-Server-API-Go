package serverClient

import (
	"client/models/responseModel"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetDollarQuote() (responseModel.DollarQuoteResponse, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return responseModel.DollarQuoteResponse{}, fmt.Errorf("erro ao criar requisição para a URL http://localhost:8080/cotacao: \n -%w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {

		if errors.Is(err, context.DeadlineExceeded) {
			return responseModel.DollarQuoteResponse{}, fmt.Errorf("tempo limite de 300ms(client) da requisição para http://localhost:8080/cotacao foi excedido")
		}

		return responseModel.DollarQuoteResponse{}, fmt.Errorf("erro ao enviar requisição para http://localhost:8080/cotacao: \n - %w", err)
	}
	defer res.Body.Close()

	json, err := io.ReadAll(res.Body)
	if err != nil {
		return responseModel.DollarQuoteResponse{}, fmt.Errorf("erro ao ler o corpo da resposta da URL http://localhost:8080/cotacao: \n - %w", err)
	}

	if res.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("resposta HTTP com status %d da URL http://localhost:8080/cotacao.\n - Detalhes: %s", res.StatusCode, json)
		return responseModel.DollarQuoteResponse{}, errors.New(msg)
	}

	fmt.Println("\n -RESPOSTA RECEBIDA -> ", string(json))

	dollarQuote := responseModel.ConvertJsonToDollarQuoteResponse(json)

	return dollarQuote, nil
}
