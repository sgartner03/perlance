package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/sgartner03/perlance/internal"
	"net/http"
	"strconv"
)

type Handler struct {
	purchases []internal.Purchase
}

func (h *Handler) getPurchases(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.purchases)
}

func (h *Handler) getPurchase(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "404: Resource not found"})
		return
	}

	index := h.findPurchase(id)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "404: Resource not found"})
		return
	}
	purchase := h.purchases[index]

	c.IndentedJSON(http.StatusOK, purchase)
}

func (h *Handler) createPurchase(c *gin.Context) {
	var purchase internal.Purchase

	err := c.BindJSON(&purchase)
	if err != nil {
		fmt.Println(err)
	} else {
		h.purchases = append(h.purchases, purchase)
		c.IndentedJSON(http.StatusCreated, purchase)
	}
}

func (h *Handler) putPurchase(c *gin.Context) {
	var purchase internal.Purchase

	err := c.BindJSON(&purchase)
	if err != nil {
		fmt.Println(err)
	} else {
		index := h.findPurchase(purchase.Id)
		if index == -1 {
			h.purchases = append(h.purchases, purchase)
		} else {
			h.purchases[index] = purchase
		}
		c.IndentedJSON(http.StatusOK, purchase)
	}
}

func (h *Handler) deletePurchase(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "404: Resource not found"})
		return
	}

	index := h.findPurchase(id)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "404: Resource not found"})
		return
	}

	purchase := h.Remove(index)

	c.IndentedJSON(http.StatusOK, purchase)
}

func trace(c *gin.Context) {
	c.Render(http.StatusOK, render.Data{
		ContentType: "http/message",
	})
}

//TODO: Implement a more efficient searching algorithm than linear
// 		-> sort + binary search?
func (h *Handler) findPurchase(id int) int {
	for index, purchase := range h.purchases {
		if purchase.Id == id {
			return index
		}
	}
	return -1
}

func (h *Handler) Remove(i int) internal.Purchase {
	var purchase internal.Purchase

	if i >= 0 || i < len(h.purchases) {
		purchase = h.purchases[i]
		h.purchases = append(h.purchases[:i], h.purchases[i+1:]...)
	}

	return purchase
}
