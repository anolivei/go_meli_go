package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

func getHello(c *gin.Context) {
	user := os.Getenv("USER")
	str := "Olá " + user
	c.JSON(200, gin.H{"mensagem": str,})
}

func main() {
	router := gin.Default()
	router.GET("ola", getHello)
	router.Run()
}
