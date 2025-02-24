package main

import (
	"acme/product/product"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/api/product-list", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, product.ProductList)
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(fmt.Sprintf("failed to start HTTP server: %s", err))
	}
}
