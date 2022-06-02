package main

//a. O repositório, serviço e handler devem ser importados e injetados
import (
	"log"

	"github.com/anolivei/go_meli_go/tree/main/go_web_iii_aula_2/ex02/cmd/server/handler"
	"github.com/anolivei/go_meli_go/tree/main/go_web_iii_aula_2/ex02/internal/products"
	"github.com/anolivei/go_meli_go/tree/main/go_web_iii_aula_2/ex02/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//b. O roteador deve ser implementado para os diferentes endpoints
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
