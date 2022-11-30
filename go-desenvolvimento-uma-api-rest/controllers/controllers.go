package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"learning-go/go-desenvolvimento-uma-api-rest/database"
	"learning-go/go-desenvolvimento-uma-api-rest/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades) // p é um array de personalidades	json.NewEncoder(w).Encode(models.Personalidades)
	json.NewEncoder(w).Encode(personalidades)
	//json.NewEncoder(w).Encode(models.Personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // pega todos os valores da req
	id := vars["id"]    // identifico que quero so o valor do id
	var personalidade models.Personalidade
	// Exemplo sem BD
	//for _, personalidade := range models.Personalidades {
	//	if strconv.Itoa(personalidade.Id) == id { // strconv.Itoa: converte o valor  inteiro para string
	//		json.NewEncoder(w).Encode(personalidade)
	//	}
	//}
	database.DB.First(&personalidade, id) // &p = busco o endereço de memoria, do atributo pelo id
	json.NewEncoder(w).Encode(personalidade)
}
