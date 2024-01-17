package rotas

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Representa as rotas da aplucação
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas no router
func Configurar(router *mux.Router) *mux.Router {

	rotas := rotasLogin
	rotas = append(rotas, RotaPaginaPrincipal)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			router.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}

	//Configura o fileserver para que ele fique no /assets/
	fileserver := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileserver))

	return router

}
