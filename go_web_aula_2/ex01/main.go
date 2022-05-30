package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func getById(usuarios []usuario)(gin.HandlerFunc) {
	fn := func(c *gin.Context) {
		id := c.Query("id")
		id2, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "id is not a number",
			})
			log.Println(err)
			return
		}
		for _, u := range usuarios {
			fmt.Println(u.Id)
			if u.Id == id2 {
				c.JSON(200, u)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"message": "id not found",
		})
	}
	return gin.HandlerFunc(fn)
}

func createMatrixUsers()([]usuario) {
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
	return u
}

func main() {
	u := createMatrixUsers()
	gin.SetMode("release")
	router := gin.Default()
	router.GET("users", getById(u))
	router.Run()
}
