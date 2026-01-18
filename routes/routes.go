package routes

import (
	"LojaDoHugo/controllers"
	"net/http"
)

func CarregaRota() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
