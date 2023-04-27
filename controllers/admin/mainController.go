package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		models.DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC")
		}).Order("sort desc").Find(&accessList)

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

		//fmt.Printf("this is %#v", accessList)
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

func (con MainController) ChangeStatus(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")

	// status = ABS(0-1)   1
	// status = ABS(1-1)  0
	err1 := models.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败 请重试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}

// 公共修改状态的方法
func (con MainController) ChangeNum(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")
	num := c.Query("num")

	err1 := models.DB.Exec("update "+table+" set "+field+"="+num+" where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改数据失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "修改成功",
		})
	}

}

// 公共修改状态的方法
func (con MainController) ChangePath(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")
	path := c.Query("path")

	err1 := models.DB.Debug().Exec("update "+table+" set "+field+"='"+path+"' where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改数据失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "修改成功",
		})
	}

}
