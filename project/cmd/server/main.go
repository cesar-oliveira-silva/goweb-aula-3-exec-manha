package main

import (
	"log"

	"github.com/cesar-oliveira-silva/goweb-aula-3-exec-manha.git/project/cmd/server/handler"
	"github.com/cesar-oliveira-silva/goweb-aula-3-exec-manha.git/project/internal/usuarios"
	"github.com/gin-gonic/gin"
)

// var dbConn *sql.DB

func main() {
	repo := usuarios.NewRepository()
	service := usuarios.NewService(repo)
	userHandler := handler.NewUser(service)

	server := gin.Default()
	pr := server.Group("/usuarios")
	pr.POST("/", userHandler.Store())
	pr.GET("/", userHandler.GetAll())
	pr.PUT("/:id", userHandler.Update())
	pr.PATCH("/:id", userHandler.UpdateName())
	pr.DELETE("/:id", userHandler.Delete())

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
