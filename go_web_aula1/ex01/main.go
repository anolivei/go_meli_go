package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type usuario struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email string `json:"email"`
	Idade int `json:"idade"`
	Altura float32 `json:"altura"`
	Ativo bool `json:"ativo"`
	DataDeCriacao string `json:"data_de_criacao"`
}

func main() {
	u := usuario{555,
				"Joao",
				"das Neves",
				"jdn@gmail.com",
				30,
				1.75,
				true,
				"28/05/2022"}

	content, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("./go-web/users.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		panic(err)
	}

	fmt.Println("users.json created!")
}