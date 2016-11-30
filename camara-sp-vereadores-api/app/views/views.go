package views

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)
import "github.com/alexandre/lab-go/camara-sp-vereadores-api/app/controllers"

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Fprintf(w, "Hello world!")
}

func Vereadores(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	vereadores := controllers.GetAllVereadores(session)

	data, _ := json.Marshal(vereadores)

	w.Write(data)
}

func Vereador(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	vereadorId := ps.ByName("id")

	vereador, err := controllers.GetVereador(session, vereadorId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		data, _ := json.Marshal(vereador)
		w.Write(data)
	}
}
