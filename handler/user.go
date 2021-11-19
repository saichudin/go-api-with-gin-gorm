package handler

import (
	"api-crud-with-gin/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewBookHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RootHandler(c *gin.Context) {
	user, err := h.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
