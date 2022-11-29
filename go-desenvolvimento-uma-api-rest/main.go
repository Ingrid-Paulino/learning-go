package main

import (
	"fmt"
	"learning-go/go-desenvolvimento-uma-api-rest/models"
	"learning-go/go-desenvolvimento-uma-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
