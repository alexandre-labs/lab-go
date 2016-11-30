package store_data

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)
import "github.com/alexandre/lab-go/camara-sp-vereadores-api/app/structs"

func main() {

	absPath, err := filepath.Abs("./data")
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		panic(err)
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("camara-sp-vereadores").C("vereadores")
	collection.DropCollection()

	for _, file := range files {

		content, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			panic(err)
		}

		var vereador structs.Vereador
		err = json.Unmarshal(content, &vereador)
		if err != nil {
			panic(err)
		}

		err = collection.Insert(vereador)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Done.")

}
