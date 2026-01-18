package main

import (
	"LojaDoHugo/routes"
	"net/http"
)

func main() {
	routes.CarregaRota()
	http.ListenAndServe(":8080", nil)
}
