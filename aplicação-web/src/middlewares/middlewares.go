package middlewares

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		// Aqui ficam as instruções que serão executadas antes da requisição chegar ao controller
		// ou seja, é um middleware
		proximaFuncao(w, r)
		// Aqui ficam as instruções que serão executadas depois da requisição chegar ao controller
		// ou seja, é um middleware
	}
}

// Verificar se os dados estão no cookie, e não se estão corretos (API que faz isso)
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !IpExiste(r.RemoteAddr) {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			fmt.Println(r.RemoteAddr)
			fmt.Println(r.Host)
			fmt.Println("IP não autorizado")
			return
		}

		//fmt.Println(valores, erro)

		// Aqui ficam as instruções que serão executadas antes da requisição chegar ao controller
		// ou seja, é um middleware
		proximaFuncao(w, r)
		// Aqui ficam as instruções que serão executadas depois da requisição chegar ao controller
		// ou seja, é um middleware
	}
}

func IpExiste(ip string) bool {
	// Divide o IP na ":" e pega a primeira parte (o IP sem a porta)
	ip = strings.Split(ip, ":")[0]

	file, err := os.Open("/var/www/html/teste/ip.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), ip) {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return false
}

func PegarNomePorIp(ip string) (string, error) {

	// Divide o IP na ":" e pega a primeira parte (o IP sem a porta)
	ip = strings.Split(ip, ":")[0]

	file, err := os.Open("/var/www/html/teste/ip.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 2 && fields[0] == ip {
			return fields[1], nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", nil
}
