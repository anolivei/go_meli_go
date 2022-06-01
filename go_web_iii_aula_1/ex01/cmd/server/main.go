package main

//a. O repositório, serviço e handler devem ser importados e injetados
import (
	"github.com/anolivei/go_meli_go/tree/main/go_web_iii_aula_1/ex01/cmd/server/handler"
	"github.com/anolivei/go_meli_go/tree/main/go_web_iii_aula_1/ex01/internal/products"
	"github.com/gin-gonic/gin"
)

//b. O roteador deve ser implementado para os diferentes endpoints
func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	r.Run()
}
