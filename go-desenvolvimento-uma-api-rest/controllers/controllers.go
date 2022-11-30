package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"learning-go/go-desenvolvimento-uma-api-rest/models"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // pega todos os valores da req
	id := vars["id"]    // identifico que quero so o valor do id

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id { // strconv.Itoa: converte o valor  inteiro para string
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
