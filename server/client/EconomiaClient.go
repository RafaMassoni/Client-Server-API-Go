package client

import (
	"context"
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
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
