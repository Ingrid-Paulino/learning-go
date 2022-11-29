package routes

import (
	"learning-go/go-desenvolvimento-uma-api-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades) //sempre que uma req. cair na rota "/" a funcao Home sera invocada
	log.Fatal(http.ListenAndServe(":8000", nil))                            //sobe o servidor
}
