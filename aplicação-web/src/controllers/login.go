package controllers

import (
	"net/http"
	"webapp/src/respostas"
)

func FazerLogin(w http.ResponseWriter, r *http.Request) {

	respostas.JSON(w, http.StatusOK, nil)

}
