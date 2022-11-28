package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {

	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: CreateHandlers(),
	}

	srv.ListenAndServe()
}

func CreateHandlers() http.Handler {
	handler := gin.Default()

	handler.GET("/users", GetAllUsers)
	handler.POST("/users", CreateNewUser)
	handler.GET("/users/:id", GetUserById)

	return handler
}

func GetAllUsers(c *gin.Context) {
	c.String(http.StatusOK, "All users")
}

func CreateNewUser(c *gin.Context) {

	userData := &UserRequest{}

	if err := json.NewDecoder(c.Request.Body).Decode(userData); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Printf("%#v\n", userData)

	c.AbortWithStatus(http.StatusNoContent)
}

func GetUserById(c *gin.Context) {
	// id, ok := c.Params.Get("id")
	// if !ok {
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// }

	userData := &UserRequest{
		ID:       3456,
		Username: "Vasiliy Pupkin",
		Email:    "email@example.org",
	}

	c.JSON(http.StatusOK, userData)
}
