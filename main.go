package main

import (
	"fmt"

	"github.com/lourencogabe/api-go/database"
	"github.com/lourencogabe/api-go/routes"
)

func main() {
	database.ConectionDataBase()
	fmt.Println("Iniciando o servidor em GO")
	routes.HandleResquest()
}
