package handlers

import (
	"fmt"
	"go-cred-app/models"
	"go-cred-app/store"
	"go-cred-app/utils"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type loginBody struct {
	Userid   string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var body loginBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var userDetails models.User

	for _, value := range store.Users {
		if value.Userid == body.Userid {
			userDetails = value
			break
		}
	}

	if userDetails.Userid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
		return
	}

	if userDetails.Password != body.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong password"})
		return
	}

	jwtToken, err := utils.GenerateJwtToken(body.Userid, userDetails.Firstname)
	if err != nil {
		fmt.Println("Token generation error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error generating token"})
		return
	}

	fmt.Println("token: ", jwtToken)

	c.JSON(http.StatusOK, gin.H{"message": "Login success", "token": jwtToken})
}

func RegisterHandler(c *gin.Context) {
	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	firstName := body.Firstname
	userExists := slices.ContainsFunc(store.Users, func(u models.User) bool {
		return u.Firstname == firstName
	})

	fmt.Println(userExists)

	if userExists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exist"})
		return
	}

	var uniqueId = uuid.New()

	body.Userid = uniqueId.String()
	store.Users = append(store.Users, body)

	c.JSON(http.StatusCreated, gin.H{"message": "Registration completed", "userId": body.Userid})

	fmt.Println("Total users: &v", store.Users)
}
