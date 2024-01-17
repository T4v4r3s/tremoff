package rotas

import "webapp/src/controllers"

var RotaPaginaPrincipal = Rota{
	URI:                "/home",
	Metodo:             "GET",
	Funcao:             controllers.CarregarPaginaPrincipal,
	RequerAutenticacao: true,
}
