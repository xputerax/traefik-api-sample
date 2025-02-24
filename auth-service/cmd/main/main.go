package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	ClientId  string `json:"client_id"`
	SecretKey string `json:"secret_key"`
}

func main() {
	router := gin.New()

	router.POST("/api/auth", func(ctx *gin.Context) {
		var req AuthRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "auth failed",
				"error":   "bad request",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "auth ok",
		})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
			"error":   fmt.Sprintf("resource %s not found", ctx.Request.URL.Path),
		})
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(fmt.Sprintf("failed to start server: %s", err))
	}
}
