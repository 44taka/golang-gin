package main

import (
	"github.com/44taka/golang-gin/infrastructure"
	"github.com/44taka/golang-gin/infrastructure/persistence"
	"github.com/44taka/golang-gin/interfaces/handler"
	"github.com/44taka/golang-gin/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// envファイル読み込み
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// コンフィグ読み込み
	config := infrastructure.NewConfig()
	db := infrastructure.NewDB(config)

	// マイグレーション実行
	err = infrastructure.Migrate(config)
	if err != nil {
		panic(err)
	}

	// ハンドラー読み込み
	userPersistence := persistence.NewUserPersistence(db.Connect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	r := gin.Default()
	r.GET("/users", func(ctx *gin.Context) { userHandler.FindAll(ctx) })
	r.GET("/users/:id", func(ctx *gin.Context) { userHandler.FindById(ctx) })
	r.POST("/users", func(ctx *gin.Context) { userHandler.Create(ctx) })
	r.PUT("/users/:id", func(ctx *gin.Context) { userHandler.Update(ctx) })
	r.DELETE("/users/:id", func(ctx *gin.Context) { userHandler.Delete(ctx) })

	r.Run(":" + config.Routing.Port)
}
