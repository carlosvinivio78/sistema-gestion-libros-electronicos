package routes

import (
	"sistema-gestion-libros-electronicos/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/books")
	{
		api.GET("/", handlers.GetBooksJSON)
		api.GET("/:id", handlers.GetBookByIDJSON)
		api.POST("/", handlers.CreateBookJSON)
		api.PUT("/:id", handlers.UpdateBookJSON)
		api.DELETE("/:id", handlers.DeleteBookJSON)
	}

	return r
}
