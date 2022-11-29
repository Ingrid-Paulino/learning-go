package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("Go-crie-uma-aplicacao-web/templates/index.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// template.Must encapsula todos os meus templates e devolve dois retornos: o template e uma mensagem de erro -  se o erro for nil==nulo sig. que ele conseguiu pegar o template
// template.ParseGlob: passo o caminho onde esta os meus temlates
//var temp = template.Must(template.ParseGlob("Go-crie-uma-aplicacao-web/templates/*.html"))
//
//func main() {
//	http.HandleFunc("/", index)
//	http.ListenAndServe(":8000", nil) // sobe o servidor
//}

//func index(w http.ResponseWriter, r *http.Request) {
//	produtos := []Produto{
//		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
//		{"Tenis", "Confortável", 89, 3},
//		{"Fone", "Muito bom", 59, 2},
//		{"Produto novo", "Muito legal", 1.99, 1},
//	}
//	//w: quem consegue escrever/passar para nossa resposta
//	//index: quem ele vai mostrar - obs: esse index esta na primeira linha do arquivo HTML
//	//nil: como nao quero passar nenhum informacao, entt passo o nil
//	//temp.ExecuteTemplate(w, "Index", nil)
//	temp.ExecuteTemplate(w, "Index", produtos)
//
//}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
		{"Bolsa", "Muito linda", 15.99, 1},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
