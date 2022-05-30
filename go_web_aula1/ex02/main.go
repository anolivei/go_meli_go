package main

import "github.com/gin-gonic/gin"

func getHello(c *gin.Context) {
	c.JSON(200, gin.H{"mensagem": "Olá Angélica",})
}

func main() {
	router := gin.Default()
	router.GET("ola", getHello)
	router.Run()
}
