package routers

import (
	"github.com/gin-gonic/gin"
	"neverGiveUp/controllers/admin"
	"neverGiveUp/middleWare"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middleWare.InitMiddleWare)
	{
		adminRouters.GET("/", admin.IndexController{}.Index)
		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/login", admin.LoginController{}.Login)
		adminRouters.GET("/captcha", admin.LoginController{}.Captcha)
		adminRouters.GET("/manager", admin.ManagerController{}.Index)

		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouters.POST("/manager/doAdd", admin.ManagerController{}.DoAdd)

		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouters.POST("/manager/doEdit", admin.ManagerController{}.DoEdit)

		//adminRouters.GET("/user/add", admin.UserController{}.Add)
		//adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		//adminRouters.GET("/user/delete", admin.UserController{}.Delete)
	}

}
