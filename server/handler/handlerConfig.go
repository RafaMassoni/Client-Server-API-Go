package handler

import (
	"net/http"
)

func InitHandlers() {

	println("---- SERVER START ----")
	println("")

	http.HandleFunc("/cotacao", getDollarQuote)
	http.ListenAndServe(":8080", nil)
}
