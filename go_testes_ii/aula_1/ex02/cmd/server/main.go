package main

import (
	"log"
	"os"

	"github.com/anolivei/go_meli_go/tree/main/go_testes_ii/aula_1/ex02/cmd/server/handler"
	"github.com/anolivei/go_meli_go/tree/main/go_testes_ii/aula_1/ex02/docs"
	"github.com/anolivei/go_meli_go/tree/main/go_testes_ii/aula_1/ex02/internal/products"
	"github.com/anolivei/go_meli_go/tree/main/go_testes_ii/aula_1/ex02/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}
