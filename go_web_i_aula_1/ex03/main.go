package main

import (
	"encoding/json"
	"log"
	"github.com/gin-gonic/gin"
)

type usuario struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email string `json:"email"`
	Idade int `json:"idade"`
	Altura float32 `json:"altura"`
	Ativo bool `json:"ativo"`
	DataCriacao string `json:"datacriacao"`
}

func getAll(c *gin.Context) {

	jsonData := `[{"Id": 555,
		"Nome": "Joao",
		"Sobrenome": "das Neves",
		"Email": "jdn@gmail.com",
		"Idade": 30,
		"Altura": 1.75,
		"Ativo": true,
		"DataCriacao": "28/05/2022"},
		{"Id": 444,
		"Nome": "Joao",
		"Sobrenome": "das Neves",
		"Email": "jdn@gmail.com",
		"Idade": 30,
		"Altura": 1.75,
		"Ativo": true,
		"DataCriacao": "28/05/2022"}]`

	var u []usuario

	if err := json.Unmarshal([]byte(jsonData), &u); err != nil {
		log.Fatal(err)
	}
	for _, i := range u {
		c.JSON(200, i)
	}
}

func main() {
	gin.SetMode("release")
	router := gin.Default()
	router.GET("users", getAll)
	router.Run()
}
