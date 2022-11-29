package models

type Personalidade struct {
	Nome     string `json:"nome"` // json.. -> forma que o nome vai ser exibido na api rest
	Historia string `json:"historia"`
}

var Personalidades []Personalidade
