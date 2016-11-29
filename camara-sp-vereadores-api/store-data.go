package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)
import "./structs"

func main() {

	absPath, err := filepath.Abs("./data")
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {

		content, err := ioutil.ReadFile(filepath.Join("./data", file.Name()))
		if err != nil {
			panic(err)
		}

		var vereador structs.Vereador
		err = json.Unmarshal(content, &vereador)
		if err != nil {
			panic(err)
		}

		fmt.Println("Nome: ", vereador.RawData.Nome)
		fmt.Println("Partido: ", vereador.RawData.Partido.Nome)
	}

}
