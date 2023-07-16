package main

import (
	"github.com/gin-gonic/gin"
	_ "sgartner/perlance/internal"
)

func main() {
	router := gin.Default()
	var handler Handler

	router.GET("/purchases", handler.getPurchases)
	router.GET("purchases/:id", handler.getPurchase)

	router.POST("/purchases", handler.createPurchase)
	router.PUT("/purchases", handler.putPurchase)
	router.DELETE("/purchases/:id", handler.deletePurchase)

	router.Handle("TRACE", "/*all")
	router.Run("localhost:8080")
}
