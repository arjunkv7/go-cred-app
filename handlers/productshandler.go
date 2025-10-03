package handlers

import (
	"go-cred-app/models"
	"go-cred-app/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	productId:= c.Query("productId")
	for _, product := range store.Products {
		if product.Productid == productId {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusOK, store.Products)
}

func AddProduct(c *gin.Context) {
	var product models.Products
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productId := product.Productid

	for _, product := range store.Products {
		if product.Productid == productId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product already exists"})
			return
		}
	}
	store.Products = append(store.Products, product)
	c.JSON(http.StatusOK, product)
}
