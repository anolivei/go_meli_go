package main

import (
	"log"

	"github.com/anolivei/go_meli_go/tree/main/go_web_iv_aula_1/ex01/cmd/server/handler"
	"github.com/anolivei/go_meli_go/tree/main/go_web_iv_aula_1/ex01/internal/products"
	"github.com/anolivei/go_meli_go/tree/main/go_web_iv_aula_1/ex01/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("failed to load .env")
	}
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}
