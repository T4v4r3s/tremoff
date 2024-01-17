package controllers

import (
	"net/http"
	"webapp/src/middlewares"
	"webapp/src/utils"
)

//Aqui ficam todas as funções que carregam páginas!

// CarregarTelaDeLogin carrega tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {

	if middlewares.IpExiste(r.Host) {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		return
	}

	utils.ExecutarTemplate(w, "login.html", nil) //Carrega a página de login não passando dados para ela!
}

// CarregarPaginaPrincipal carrega a página principal
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {

	utils.ExecutarTemplate(w, "home.html", nil)
}
