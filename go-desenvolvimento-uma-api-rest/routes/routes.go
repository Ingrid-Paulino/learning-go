package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"learning-go/go-desenvolvimento-uma-api-rest/controllers"
	"learning-go/go-desenvolvimento-uma-api-rest/middleware"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()                    //crio uma instancia do gorilla/mux -- cria uma nova rota/mapeamento
	r.Use(middleware.ContentTypeMiddleware) // middleware: roda Antes de entrar em cada controller, essa linha sera executada

	//http.HandleFunc("/", controllers.Home)
	//http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades) //sempre que uma req. cair na rota "/" a funcao Home sera invocada
	//log.Fatal(http.ListenAndServe(":8000", nil)) //sobe o servidor

	// Como vou usar o gorilla/mux nao cou mais precisar do http
	r.HandleFunc("/", controllers.Home).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")

	//log.Fatal(http.ListenAndServe(":8000", r))

	// cors permite a comunicacao e estabelece seguranca na comunicaçao entre o frontend com o backend
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))) // permite que qualquer aplicacao consegue acessar e consumir essa api
	//log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins(string{"localhost:3000"}))))
}
