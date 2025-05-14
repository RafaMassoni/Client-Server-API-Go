package main

import (
	"server/handler"
	"server/services/database"
)

func main() {
	println("")
	println("O servidor está em processo de inicialização")
	database.InitDataBase()
	handler.InitHandlers()
}
