package controllers

import (
	"gopkg.in/mgo.v2"
)
import "../structs"

func GetAllVereadores(session *mgo.Session) structs.Vereadores {

	var vereadores structs.Vereadores

	collection := session.DB("camara-sp-vereadores").C("vereadores")

	err := collection.Find(nil).All(&vereadores)
	if err != nil {
		panic(err)
	}

	return vereadores
}
