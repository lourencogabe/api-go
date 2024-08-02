package main

import (
	"fmt"

	"github.com/lourencogabe/api-go/models"
	"github.com/lourencogabe/api-go/routes"
)

func main() {
	models.Users = []models.User{
		{
			Id:       1,
			Name:     "Ana Pereira",
			Age:      25,
			Document: "123.456.789-00",
			Address:  "Rua das Flores, 123, São Paulo, SP",
		},
		{
			Id:       2,
			Name:     "Carlos Lima",
			Age:      40,
			Document: "654.321.987-00",
			Address:  "Avenida das Nações, 202, Brasília, DF",
		},
	}

	fmt.Println("Iniciando o servidor em GO")
	routes.HandleResquest()
}
