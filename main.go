package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"neverGiveUp/models"
	"neverGiveUp/routers"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"Str2Html":   models.Str2Html,
		"FormatImg":  models.FormatImg,
		"Sub":        models.Sub,
		"Substr":     models.Substr,
	})
	r.LoadHTMLGlob("templates/**/**/*")
	r.Static("/static", "./static")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mySession", store))

	routers.AdminRoutersInit(r)
	routers.DefaultRoutersInit(r)

	r.Run(":9999")
}
