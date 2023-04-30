package middleWare

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"neverGiveUp/models"
	"os"
	"strings"
)

func InitAuthMiddleWare(c *gin.Context) {
	pathname := c.Request.URL.String()
	splitPathname := strings.Split(pathname, "?")[0]
	//拿到路径url
	session := sessions.Default(c)
	userInfo := session.Get("userinfo")
	userInfoStr, ok := userInfo.(string)
	fmt.Printf("userInfo = %v\n", userInfoStr)
	if ok {
		var userInfoStruct []models.Manager
		//将string 断言后的userInfo信息 传入userInfoStruct 切片
		err := json.Unmarshal([]byte(userInfoStr), &userInfoStruct)
		if err != nil || !(len(userInfoStruct) > 0 && userInfoStruct[0].Username != "") {
			//如果不是以下路径进行跳转
			if splitPathname != "/admin/login" && splitPathname != "/admin/doLogin" && splitPathname != "/admin/captcha" {
				fmt.Println("进行了一次重定向")
				c.Redirect(302, "/admin/login")
			}
		} else {
			//权限 验证成功
			fmt.Println("登录成功 进行权限判断")
			//根据路径url 获取权限id
			urlPath := strings.Replace(splitPathname, "/admin/", "", 1)
			fmt.Printf("urlPath :%s\n", urlPath)
			//如果不是是超级管理员 且 路径不是一些排除在路径 进行权限判断
			if userInfoStruct[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) {
				// 1、根据角色获取当前角色的权限列表,然后把权限id放在一个map类型的对象里面
				roleAccess := []models.RoleAccess{}
				//根据角色获得该角色的权限列表
				models.DB.Where("role_id=?", userInfoStruct[0].RoleId).Find(&roleAccess)
				roleAccessMap := make(map[int]int)
				for _, v := range roleAccess {
					roleAccessMap[v.AccessId] = v.AccessId
				}
				// 2、获取当前访问的url对应的权限id 判断权限id是否在角色对应的权限
				// pathname      /admin/manager
				access := models.Access{}
				models.DB.Where("url=?", urlPath).Find(&access)
				//3、判断当前访问的url对应的权限id 是否在权限列表的id中
				if _, ok := roleAccessMap[access.Id]; !ok {
					c.String(200, "没有权限")
					c.Abort()
				}
			}
		}
	} else {
		if splitPathname != "/admin/login" && splitPathname != "/admin/doLogin" && splitPathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}

}

//排除权限判断的方法

func excludeAuthPath(urlPath string) bool {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()

	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")
	// return true
	fmt.Println(excludeAuthPathSlice)
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
