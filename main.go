package main

import (
	"github.com/eduzera0/go-gin-api/database"
	"github.com/eduzera0/go-gin-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
