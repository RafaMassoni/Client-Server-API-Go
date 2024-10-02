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
		return responseModel.DollarQuoteResponse{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {

		fmt.Println("REQUEST ERROR ")

		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Request timed out")
		}

		return responseModel.DollarQuoteResponse{}, err
	}
	defer res.Body.Close()

	json, err := io.ReadAll(res.Body)
	if err != nil {
		return responseModel.DollarQuoteResponse{}, err
	}

	if res.StatusCode != 200 {
		msg := fmt.Sprintf("\nERROR -> %s  %s", res.Status, json)
		return responseModel.DollarQuoteResponse{}, errors.New(msg)

	}

	fmt.Println("\n   RESPONSE ->", string(json))

	dollarQuote := responseModel.ConvertJsonToDollarQuoteResponse(json)

	return dollarQuote, nil
}
