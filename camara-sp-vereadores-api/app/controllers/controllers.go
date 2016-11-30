package controllers

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
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

func GetVereador(session *mgo.Session, vereadorId string) (v structs.Vereador, err error) {

	collection := session.DB("camara-sp-vereadores").C("vereadores")

	vereador := structs.Vereador{}

	err = collection.Find(bson.M{"id": vereadorId}).One(&vereador)
	if err != nil {
		log.Print(err)
		log.Print("vereadorId: ", vereadorId)
		return vereador, err
	}

	return vereador, err
}
