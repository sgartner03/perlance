package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sgartner03/perlance/internal"
)

func main() {
	router := gin.Default()
	var handler Handler

	router.GET("/purchases", handler.getPurchases)
	router.GET("purchases/:id", handler.getPurchase)

	router.POST("/purchases", handler.createPurchase)
	router.PUT("/purchases", handler.putPurchase)
	router.DELETE("/purchases/:id", handler.deletePurchase)

	router.NoRoute(trace)
	router.Run(":8080")
}
