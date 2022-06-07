package handler

import (
	"net/http"
	"strconv"

	"github.com/anolivei/go_meli_go/tree/main/go_web_iii/aula_1/ex01/internal/products"
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

func (prod *Product) Update() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "bananasdepijama" {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "id inválido"})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			c.JSON(400, gin.H{"error": "o nome do produto é obrigatório"})
			return
		}
		if req.Typee == "" {
			c.JSON(400, gin.H{"error": "o tipo do produto é obrigatório"})
			return
		}
		if req.Count == 0 {
			c.JSON(400, gin.H{"error": "a quantidade do produto é obrigatória"})
			return
		}
		if req.Price == 0 {
			c.JSON(400, gin.H{"error": "o preço do produto é obrigatório"})
			return
		}
		p, err := prod.service.Update(int(id), req.Name, req.Typee, req.Count, req.Price)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, p)
	}
	return fn
}
