package main

import (
	"log"

	"github.com/cesar-oliveira-silva/goweb-aula-2.git/project/cmd/server/handler"
	"github.com/cesar-oliveira-silva/goweb-aula-2.git/project/internal/usuarios"
	"github.com/gin-gonic/gin"
)

// var dbConn *sql.DB

func main() {
	repo := usuarios.NewRepository()
	service := usuarios.NewService(repo)
	productHandler := handler.NewUser(service)

	server := gin.Default()
	pr := server.Group("/usuarios")
	pr.POST("/", productHandler.Store())
	pr.GET("/", productHandler.GetAll())
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
