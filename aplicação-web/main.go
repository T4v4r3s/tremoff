package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"

	"github.com/gorilla/handlers"
)

//Cria uma hash e um block key para o cookie
/* func init() {
	hashkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashkey)

	blockkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(blockkey)
} */

func main() {
	config.Carregar()
	r := router.Gerar()
	utils.CarregarTemplates() //pode ser feito numa função init também

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Printf("Rodando WebApp! Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
