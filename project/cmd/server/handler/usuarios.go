package handler

import (
	"fmt"
	"net/http"

	"github.com/cesar-oliveira-silva/goweb-aula-2.git/project/internal/usuarios"
	"github.com/gin-gonic/gin"
)

type CreateRequestDto struct {
	Nome        string `json:"nome"`
	Sobrenome   string `json:"sobrenome"`
	Email       string `json:"email"`
	Idade       int    `json:"idade"`
	Altura      int    `json:"altura"`
	Ativo       bool   `json:"ativo"`
	DataCriacao string `json:"dataCriacao"`
}

type ProductHandler struct {
	service usuarios.Service
}

func NewUser(p usuarios.Service) *ProductHandler {
	return &ProductHandler{
		service: p,
	}
}

func (c *ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			// status StatusUnauthorized equivalente ao 401
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if len(p) == 0 {
			ctx.Status(http.StatusNoContent)
			return
		}

		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		var req CreateRequestDto
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		// quando chamamos a service, os dados já estarão tratados
		fmt.Println(req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.DataCriacao)
		p, err := c.service.Store(req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.DataCriacao)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, p)
	}
}
