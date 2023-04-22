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
		//adminRouters.GET("/user/add", admin.UserController{}.Add)
		//adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		//adminRouters.GET("/user/delete", admin.UserController{}.Delete)
	}

}
