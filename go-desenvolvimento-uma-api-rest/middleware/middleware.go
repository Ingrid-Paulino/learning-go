package middleware

import "net/http"

// essa funcao seta o cabecalho
// Handler seria as funcs dos meus controllers
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json") // indica que o tipo da resposta que tem que ser retornado é um json, sem essa linha a res. nao vem certinha
		next.ServeHTTP(w, r)                               //Vai para o proximo middleware
	})

}
