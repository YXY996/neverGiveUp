package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"neverGiveUp/models"
)

type MainController struct {
	BaseController
}

func (con MainController) Index(c *gin.Context) {
	session := sessions.Default(c)
	userInfo := session.Get("userinfo")
	userInfoStr, ok := userInfo.(string)
	if ok {
		var userInfoStruct []models.Manager
		json.Unmarshal([]byte(userInfoStr), &userInfoStruct)
		//获得所有权限
		accessList := []models.Access{}
		models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

		fmt.Printf("accessList %#v\n", accessList)
		//保存权限到map
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id=?", userInfoStruct[0].RoleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}

		//执行map循环 将checked标记
		for i := 0; i < len(accessList); i++ {
			if _, ok := roleAccessMap[accessList[i].Id]; ok {
				accessList[i].Checked = true
			}
			for j := 0; j < len(accessList[i].AccessItem); j++ {
				if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
					accessList[i].AccessItem[j].Checked = true
				}
			}
		}

		fmt.Printf("this is %#v", accessList)
		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   userInfoStruct[0].Username,
			"accessList": accessList,
			"isSuper":    userInfoStruct[0].IsSuper,
		})
	} else {
		c.Redirect(302, "/admin/login")
	}
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
