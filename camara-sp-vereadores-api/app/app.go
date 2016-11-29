package main

import (
	"log"
	"net/http"
)

import "./views"

func main() {

	http.HandleFunc("/", views.Home)
	http.HandleFunc("/vereadores", views.Vereadores)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
