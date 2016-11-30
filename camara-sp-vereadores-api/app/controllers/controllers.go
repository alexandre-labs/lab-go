package controllers

import (
	"gopkg.in/mgo.v2"
)
import "github.com/alexandre/lab-go/camara-sp-vereadores-api/app/structs"

func GetAllVereadores(session *mgo.Session) []structs.Vereadores {

	vereadores := make([]structs.Vereadores, 1, 56)

	collection := session.DB("camara-sp-vereadores").C("vereadores")

	err := collection.Find(nil).All(&vereadores)
	if err != nil {
		panic(err)
	}

	return vereadores
}
