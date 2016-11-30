package views

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"net/http"
)
import "github.com/alexandre/lab-go/camara-sp-vereadores-api/app/controllers"

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello world!")
}

func Vereadores(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	vereadores := controllers.GetAllVereadores(session)

	data, _ := json.Marshal(vereadores)

	w.Write(data)
}
