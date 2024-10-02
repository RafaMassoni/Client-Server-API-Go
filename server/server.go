package main

import (
	"server/handler"
	"server/services/database"
)

func main() {

	database.InitDataBase()
	handler.InitHandlers()

}
