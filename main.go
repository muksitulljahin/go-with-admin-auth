package main

import (
	"go_backend/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	apiVersion1 = "/api/v1"
	port        = ":8080"
)

func main() {
	r := gin.Default()
	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "api work successfull",
		})
	})

	api := r.Group(apiVersion1)

	auth.RegisterRoutes(api)

	r.Run(port)
}
