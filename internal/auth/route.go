package auth

import "github.com/gin-gonic/gin"

// RegisterRoutes auth মডিউলের সব রাউট রেজিস্টার করে
func RegisterRoutes(router *gin.RouterGroup) {
	// ডিপেন্ডেন্সি ইনজেকশন (Service -> Handler)
	service := NewAuthService()
	handler := NewAuthHandler(service)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", handler.Login)
	}
}
