package handler

import (
	"net/http"
)

func InitHandlers() {

	println("A API está inicializando")

	http.HandleFunc("/cotacao", getDollarQuote)

	println("A API foi inicializada com sucesso")
	println("-GET -> http://localhost:8080/cotacao  disponível para consulta")
	println("")

	println("O servidor foi iniciado com sucesso")
	println("#####################################")
	println("")

	http.ListenAndServe(":8080", nil)
}
