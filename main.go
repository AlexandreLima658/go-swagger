package main

import (
	_ "github.com/AlexandreLima658/swagger-go/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// User represents a user object
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = []User{
	{ID: "1", Username: "john_doe", Email: "john@example.com"},
	{ID: "2", Username: "jane_smith", Email: "jane@example.com"},
}

// @title Documentation Progete API
// @description This is a simple API for managing users
// @host localhost:8080
// @BasePath /v1
func main() {
	r := gin.Default()

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{
		v1.GET("/users", getUsersHandler)
		v1.GET("/user/:id", getUserByIDHandler)
	}

	r.Run(":8080")
}

// @Summary Get a list of users
// @Description Get a list of all users
// @Produce json
// @Success 200 {array} User
// @Router /v1/users [get]
func getUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// @Summary Get user by ID
// @Description Get user information by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /v1/user/{id} [get]
func getUserByIDHandler(c *gin.Context) {
	id := c.Param("id")
	// Convert id to int
	// ...
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
