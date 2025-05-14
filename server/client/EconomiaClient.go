package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"server/model/responseModel"
	"time"
)

func GetDollarQuote() (responseModel.DollarQuoteResponse, error) {

	res, err := getDollarQuoteRequest()
	if err != nil {
		return responseModel.DollarQuoteResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return responseModel.DollarQuoteResponse{}, err
	}

	resEconomia := responseModel.ConvertJsonToDollarQuoteResponse(body)
	resModel := responseModel.ConvertEconomiaResponseToDollarQuoteResponse(resEconomia)

	return resModel, nil
}
func getDollarQuoteRequest() (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar requisição para API de cotação: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("timeout: limite de 200ms excedido ao requisitar cotação em https://economia.awesomeapi.com.br/json/last/USD-BRL")
		}
		if errors.Is(err, context.Canceled) {
			return nil, fmt.Errorf("requisição para API de cotação foi cancelada")
		}
		return nil, fmt.Errorf("erro inesperado ao enviar requisição para API de cotação: %w", err)
	}

	return res, nil
}
