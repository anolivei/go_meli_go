package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email string `json:"email"`
	Idade int `json:"idade"`
	Altura float32 `json:"altura"`
	Ativo bool `json:"ativo"`
	DataCriacao string `json:"datacriacao"`
}

func listUsers(users *[]user) (gin.HandlerFunc) {
	fn := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": &users})
	}
	return fn
}

func saveUsers(users *[]user, id int) (gin.HandlerFunc) {
	fn := func(c *gin.Context) {
		var u user
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H {"error": err.Error(),})
			return
		}
		id++
		u.Id = id
		*users = append(*users, u)
		c.JSON(http.StatusOK, gin.H{"data": u,})
	}
	return fn
}

func main() {
	var users []user
	var id int = 0
	router := gin.Default()
	group := router.Group("users")
	{
		group.GET("/", listUsers(&users))
		group.POST("/", saveUsers(&users, id))
	}
	router.Run()
}
