package routers

import (
	"github.com/gin-gonic/gin"
	"neverGiveUp/controllers/itying"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		//defaultRouters.GET("/", defaultTool.DefaultController{}.Index)
		//defaultRouters.GET("/thumbnail1", defaultTool.DefaultController{}.ThumbNail1)
		//defaultRouters.GET("/thumbnail2", defaultTool.DefaultController{}.ThumbNail2)
		//defaultRouters.GET("/qrcode1", defaultTool.DefaultController{}.Qrcode1)
		//defaultRouters.GET("/qrcode2", defaultTool.DefaultController{}.Qrcode2)
		defaultRouters.GET("/", itying.DefaultController{}.Index)
	}
}
