package main

import (
	"fmt"
	"learning-go/go-desenvolvimento-uma-api-rest/database"
	"learning-go/go-desenvolvimento-uma-api-rest/routes"
)

func main() {
	// Exemplo sem BD
	//models.Personalidades = []models.Personalidade{
	//	{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
	//	{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	//}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
