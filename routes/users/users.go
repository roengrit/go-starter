// users.go
package routes

import (
	"gm-startd/entities"
	"gm-startd/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @description  Get All Users
// @id UsersHandler
// @produce json
// @response 200 {object} entities.Users
// @response 500 {object} entities.ErrorResponse
// @router /users [get]
func getUsersHandler(userUC usecases.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userUC.GetAll()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"code":  entities.ErrorCodeInternalServerError,
			})
			return
		}
		c.IndentedJSON(http.StatusOK, users)
	}
}

func getUserHandler(userUC usecases.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		user, err := userUC.GetByID(id)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
			return
		}
		if user == nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.IndentedJSON(http.StatusOK, user)
	}
}

func createUserHandler(userUC usecases.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entities.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
			return
		}
		if err := userUC.CreateNew(&user); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		c.Status(http.StatusCreated)
	}
}

func deleteUserHandler(userUC usecases.UserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := userUC.DeleteByID(id); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}
		c.Status(http.StatusOK)
	}
}

func Router(userUC usecases.UserUseCase, router *gin.Engine) *gin.Engine {
	router.GET("/users", getUsersHandler(userUC))
	router.GET("/users/:id", getUserHandler(userUC))
	router.POST("/users", createUserHandler(userUC))
	router.DELETE("/users/:id", deleteUserHandler(userUC))
	return router
}
