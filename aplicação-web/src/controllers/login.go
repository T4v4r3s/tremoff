package controllers

import (
	"net/http"
	"webapp/src/middlewares"
	"webapp/src/respostas"
)

func FazerLogin(w http.ResponseWriter, r *http.Request) {

	if middlewares.IpExiste(r.Host) {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)

}
