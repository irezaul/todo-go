package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irezaul/todo-go/handlers"
)

func Route() {
	server := gin.Default()

	server.LoadHTMLGlob("templates/*.tmpl")

	frontend := server.Group("/")
	{
		frontend.GET("/", handlers.HomeFrontend)
		frontend.GET("/register", handlers.Register)
		// frontend.GET("/todo", handlers.Todo)
	}
	server.Run(":8080")
}
