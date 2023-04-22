package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	username, _ := c.Get("username")
	fmt.Printf("this is %v\n", username)
	v, ok := username.(string)
	if ok {
		c.String(http.StatusOK, "用户列表--"+v)
	} else {
		c.String(http.StatusBadRequest, "获取用户列表失败")
	}
}
