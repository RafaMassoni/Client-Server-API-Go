package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/client"
	"server/model/tableModel"
	"server/services/database"
)

func getDollarQuote(w http.ResponseWriter, r *http.Request) {

	response, err := client.GetDollarQuote()
	if err != nil {
		log.Println("ECONOMIA ERROR ", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Ops, tivemos um problema: ECONOMIA ERROR ->  %s", err.Error())
	} else {

		log.Println("DollarValue ", response.DollarValue)
		database.InsertDollarQuote(tableModel.DollarQuote{
			DollarValue: response.DollarValue,
		})

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}

}
