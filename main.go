package main

import (
	"go_backend/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiVersion1 = "/api/v1"

func main() {
	r := gin.Default()
	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "api work successfully",
		})
	})

	api := r.Group(apiVersion1)

	auth.RegisterRoutes(api)

	r.Run("localhost:8080")
}
