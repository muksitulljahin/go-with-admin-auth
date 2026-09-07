package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

// NewAuthHandler কনস্ট্রাকটর
func NewAuthHandler(service Service) *Handler {
	return &Handler{service: service}
}

// Login Controller
func (h *Handler) Login(ctx *gin.Context) {
	var req LoginRequest

	// ১. রিকোয়েস্ট বডি ভ্যালিডেশন
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ২. সার্ভিস লেয়ার কল করা
	token, err := h.service.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// ৩. সফল রেসপন্স রিটার্ন
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
