package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// コンフィグ読み込みaa
	// config := infrastructure.NewConfig()
	// db := infrastructure.NewDB(config)

	// userPersistence := persistence.NewUserPersistence(db.Connect())
	// userUseCase := usecase.NewUserUseCase(userPersistence)
	// userHandler := handler.NewUserHandler(userUseCase)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	// r.GET("/users", func(ctx *gin.Context) { userHandler.FindAll(ctx) })
	// r.GET("/users/:id", func(ctx *gin.Context) { userHandler.FindById(ctx) })
	// r.POST("/users", func(ctx *gin.Context) { userHandler.Create(ctx) })
	// r.PUT("/users/:id", func(ctx *gin.Context) { userHandler.Update(ctx) })
	// r.DELETE("/users/:id", func(ctx *gin.Context) { userHandler.Delete(ctx) })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
