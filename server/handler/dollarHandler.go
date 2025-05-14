package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"server/client"
	"server/model/tableModel"
	"server/services/database"
)

func getDollarQuote(w http.ResponseWriter, r *http.Request) {

	log.Printf("GET -> /cotacao | uma request foi identificada")

	response, err := client.GetDollarQuote()
	if err != nil {

		if errors.Is(err, context.Canceled) && r.Context().Err() == context.Canceled {
			log.Println("ECONOMIA ERROR: requisição cancelada pelo client \n-", err)
			http.Error(w, fmt.Sprintf("ECONOMIA ERROR: requisição cancelada pelo client \n-%v", err), http.StatusInternalServerError)

		} else {
			log.Println("Um erro ocorreu ao consultar o serviço Economia:", err)
			http.Error(w, fmt.Sprintf("Um erro ocorreu ao consultar o serviço Economia: -%v", err), http.StatusInternalServerError)
		}

	} else {

		log.Println("DollarValue ", response.DollarValue)
		insertErr := database.InsertDollarQuote(tableModel.DollarQuote{
			DollarValue: response.DollarValue,
		})

		if insertErr != nil {

			if errors.Is(insertErr, context.Canceled) && r.Context().Err() == context.Canceled {
				log.Println("DATABASE ERROR: requisição cancelada pelo client \n-", insertErr)
				http.Error(w, fmt.Sprintf("DATABASE ERROR: requisição cancelada pelo client \n-%v", insertErr), http.StatusInternalServerError)
			} else {
				log.Println("Um erro ocorreu ao registrar a cotação no banco de dados:", insertErr)
				http.Error(w, fmt.Sprintf("Um erro ocorreu ao registrar a cotação no banco de dados: -%v", insertErr), http.StatusInternalServerError)
			}

		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}

}
