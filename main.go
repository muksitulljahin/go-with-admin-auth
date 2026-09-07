package main

import (
	"go_backend/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// docs ফোল্ডার জেনারেট হওয়ার পর এটি লাগবে
)

// @title           Go Admin & Auth API
// @version         1.0
// @description     This is a sample authentication server.
// @host            localhost:8080
// @BasePath        /api/v1

const (
	apiVersion1 = "/api/v1"
	port        = ":8080"
)

func main() {
	r := gin.Default()

	// Swagger UI Route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "api work successfull",
		})
	})

	api := r.Group(apiVersion1)
	auth.RegisterRoutes(api)

	r.Run(port)
}
