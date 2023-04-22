package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"neverGiveUp/models"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	userList := []models.User{}
	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
}
