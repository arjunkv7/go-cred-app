package handlers

import (
	"context"
	"fmt"
	//"fmt"
	"go-cred-app/config"
	"go-cred-app/models"
	"go-cred-app/store"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

func GetProductFromDB(c *gin.Context) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	productId := c.Query("productId")
	var products []models.Products
	fmt.Println("productId: ", productId)
	if productId != "" {
		var product models.Products
		result := config.DbClient.Collection("products").FindOne(context, bson.M{"productid": productId})
		result.Decode(&product)
		products = append(products, product)
		
	} else {
		// var product models.Products
		cursor, err := config.DbClient.Collection("products").Find(context, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(context)

		err = cursor.All(context, &products)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// c.JSON(http.StatusOK, products)
		fmt.Println(products)
	}
	
	c.JSON(http.StatusOK, products)
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

func AddProductToDB(c *gin.Context) {
	var product models.Products
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var productExists models.Products

	err := config.DbClient.Collection("products").FindOne(context, bson.M{"productid": product.Productid}).Decode(&productExists)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product already exists"})
		return
	}

	config.DbClient.Collection("products").InsertOne(context, product)
	c.JSON(http.StatusOK, gin.H{"message": "Product added to database"})
}
