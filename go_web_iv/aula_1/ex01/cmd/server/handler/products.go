package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/anolivei/go_meli_go/tree/main/go_web_iv/aula_1/ex01/internal/products"
	"github.com/anolivei/go_meli_go/tree/main/go_web_iv/aula_1/ex01/pkg/web"
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

type request struct {
	Name string `json:"name"`
	Typee string `json:"type"`
	Count int `json:"count"`
	Price float64 `json:"price"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{service: p}
}

func (prod *Product) GetAll() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN"){
			c.JSON(http.StatusUnauthorized,
				web.NewResponse(http.StatusUnauthorized, nil, ERROR_TOKEN))
			return
		}
		p, err := prod.service.GetAll()
		if err != nil {
			c.JSON(http.StatusNotFound,
				web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
	return fn
}

func (prod *Product) Store() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized,
				web.NewResponse(http.StatusUnauthorized, nil, ERROR_TOKEN))
			return
		}
		var req request
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusNotFound,
				web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		p, err := prod.service.Store(req.Name, req.Typee, req.Count, req.Price)
		if err != nil {
			c.JSON(http.StatusNotFound,
				web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
	return fn
}

func (prod *Product) Update() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized,
				web.NewResponse(http.StatusUnauthorized, nil, ERROR_TOKEN))
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, ERROR_ID))
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if req.Name == "" {
			c.JSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, ERROR_NAME))
			return
		}
		if req.Typee == "" {
			c.JSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, ERROR_TYPE))
			return
		}
		if req.Count == 0 {
			c.JSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, ERROR_COUNT))
			return
		}
		if req.Price == 0 {
			c.JSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil,  ERROR_PRICE))
			return
		}
		p, err := prod.service.Update(int(id),
									req.Name,
									req.Typee,
									req.Count,
									req.Price)
		if err != nil {
			c.JSON(http.StatusNotFound,
				web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
	return fn
}

func (prod *Product) UpdateName() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, 
				web.NewResponse(http.StatusUnauthorized, nil, ERROR_TOKEN))
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, ERROR_ID))
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if req.Name == "" {
			c.JSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, ERROR_NAME))
			return
		}
		p, err := prod.service.UpdateName(int(id), req.Name)
		if err != nil {
			c.JSON(http.StatusNotFound,
				web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
	return fn
}

func (prod *Product) Delete() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized,
				web.NewResponse(http.StatusUnauthorized, nil, ERROR_TOKEN))
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, ERROR_ID))
			return
		}
		err = prod.service.Delete(int(id))
		if err != nil {
			c.JSON(http.StatusNotFound,
				web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		p := fmt.Sprintf("o produto %d foi removido", id)
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
	return fn
}
