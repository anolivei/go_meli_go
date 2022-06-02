package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/anolivei/go_meli_go/tree/main/go_web_iii_aula_1/ex03/internal/products"
	"github.com/gin-gonic/gin"
)

const (
	ERROR_NAME = "o nome do produto é obrigatório"
	ERROR_TYPE = "o tipo do produto é obrigatório"
	ERROR_COUNT = "a quantidade do produto é obrigatória"
	ERROR_PRICE = "o preço do produto é obrigatório"
	ERROR_TOKEN = "token inválido"
	ERROR_ID = "id inválido"
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
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": ERROR_NAME})
			return
		}
		if req.Typee == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": ERROR_TYPE})
			return
		}
		if req.Count == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": ERROR_COUNT})
			return
		}
		if req.Price == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": ERROR_PRICE})
			return
		}
		p, err := prod.service.Update(int(id),
									req.Name,
									req.Typee,
									req.Count,
									req.Price)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	}
	return fn
}

func (prod *Product) UpdateName() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "bananasdepijama" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR_TOKEN})
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": ERROR_ID})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": ERROR_NAME})
			return
		}
		p, err := prod.service.UpdateName(int(id), req.Name)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	}
	return fn
}

func (prod *Product) Delete() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "bananasdepijama" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR_TOKEN})
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": ERROR_ID})
			return
		}
		err = prod.service.Delete(int(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": fmt.Sprintf("o produto %d foi removido", id)})
	}
	return fn
}
