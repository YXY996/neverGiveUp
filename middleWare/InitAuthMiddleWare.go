package middleWare

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"neverGiveUp/models"
	"strings"
)

func InitAuth(c *gin.Context) {
	pathname := c.Request.URL.String()
	splitPathname := strings.Split(pathname, "?")[0]
	//拿到路径url
	session := sessions.Default(c)
	userInfo := session.Get("username")
	userInfoStr, ok := userInfo.(string)
	if ok {
		var userInfoStruct []models.Manager
		//将string 断言后的userInfo信息 传入userInfoStruct 切片
		err := json.Unmarshal([]byte(userInfoStr), &userInfoStruct)
		if err != nil || !(len(userInfoStruct) > 0 && userInfoStruct[0].Username != "") {
			if splitPathname != "/admin/login" || splitPathname != "/admin/doLogin" || splitPathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login")
			}
		} else {
			//权限 验证成功

		}

	} else {
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}

}
