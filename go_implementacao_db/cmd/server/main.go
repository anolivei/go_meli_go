package main

import (
	"log"
	"os"

	"github.com/anolivei/go_meli_go/tree/main/go_implementacao_db/cmd/server/handler"
	"github.com/anolivei/go_meli_go/tree/main/go_implementacao_db/docs"
	"github.com/anolivei/go_meli_go/tree/main/go_implementacao_db/internal/products"
	//"github.com/anolivei/go_meli_go/tree/main/go_implementacao_db/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
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
	//db := store.New(store.FileType, "./products.json")
	dataSource := os.Getenv("DATA_SOURCE")
	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("failed to connect to mariadb")
	}
	repo := products.NewMariaDBRepository(conn)
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
