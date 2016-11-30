package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

import "github.com/alexandre/lab-go/camara-sp-vereadores-api/app/views"

func main() {

	router := httprouter.New()
	router.GET("/", views.Home)
	router.GET("/vereadores", views.Vereadores)
	router.GET("/vereador/:id", views.Vereador)

	log.Fatal(http.ListenAndServe(":1234", router))
}
