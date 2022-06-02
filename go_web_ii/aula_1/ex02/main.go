package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type user struct {
	Id int `json:"id"`
	Nome string `json:"nome" validate:"required"`
	Sobrenome string `json:"sobrenome" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Idade int `json:"idade" validate:"required,gte=0,lt=150"`
	Altura float32 `json:"altura" validate:"required"`
	Ativo bool `json:"ativo" validate:"required"`
	DataCriacao string `json:"datacriacao" validate:"required"`
}

func listUsers(users *[]user) (gin.HandlerFunc) {
	fn := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": &users})
	}
	return fn
}

func saveUsers(users *[]user, id int) (gin.HandlerFunc) {
	fn := func(c *gin.Context) {
		var validate *validator.Validate = validator.New()
		var u user
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H {"error": err.Error(),})
			return
		}
		err := validate.Struct(u)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				fmt.Println(err)
				return
			}
			for _, err := range err.(validator.ValidationErrors) {
				if err != nil {
					s := fmt.Sprintf("campo %s é obrigatório", err.Field())
					log.Println(s)
					c.JSON(http.StatusBadRequest, gin.H{"error": s,})
					return
				}
			}
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
