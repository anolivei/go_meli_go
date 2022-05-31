package handler

import (
	"net/http"

	"github.com/anolivei/go_meli_go/tree/main/go_web_ii_aula_2/ex02/internal/products"
	"github.com/gin-gonic/gin"
)

//a. A estrutura request deve ser gerada
type request struct {
	Name	string	`json:"name"`
	Typee	string	`json:"type"`
	Count	int		`json:"count"`
	Price	float64	`json:"price"`
}

//b. A estrutura do controller que tem o serviço como campo deve ser gerada
type Product struct {
	service products.Service
}

//c. A função que retorna o controlador deve ser gerada
func NewProduct(p products.Service) *Product {
	return &Product{service: p}
}

//d. Todos os métodos correspondentes aos endpoints devem ser gerados
func (prod *Product) GetAll() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "bananasdepijama" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		p, err := prod.service.GetAll()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	}
	return fn
}

func (prod *Product) Store() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "bananasdepijama" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		p, err := prod.service.Store(req.Name, req.Typee, req.Count, req.Price)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	}
	return fn
}
