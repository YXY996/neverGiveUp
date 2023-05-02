package defaultTool

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/hunterhug/go_image"
	qrcode "github.com/skip2/go-qrcode"
	"io/ioutil"
	"neverGiveUp/controllers/admin"
	"neverGiveUp/models"
)

type DefaultController struct {
	admin.BaseController
}

func (con DefaultController) Index(c *gin.Context) {
	fmt.Println("-----------")
	fmt.Println(models.GetSettingFromColumn("SiteTitle"))
	c.String(200, "首页")
}

// 缩略图
func (con DefaultController) ThumbNail1(c *gin.Context) {
	fileName := "static/upload/0.png"
	savePath := "static/upload/0_600.png"
	err := ScaleF2F(fileName, savePath, 600)
	if err != nil {
		c.String(200, "生成图片失败")
		return
	}
	c.String(200, "生成图片成功")
}

// 缩略图
func (con DefaultController) ThumbNail2(c *gin.Context) {
	filename := "static/upload/tao.jpg"
	savepath := "static/upload/tao_400.png"
	err := ThumbnailF2F(filename, savepath, 200, 300)

	if err != nil {
		c.String(200, "生成图片失败")
		return
	}
	c.String(200, "生成图片成功")
}

func (con DefaultController) Qrcode1(c *gin.Context) {
	png, err := qrcode.Encode("www.baidu.com", qrcode.Medium, 256)
	if err != nil {
		c.String(200, "生成失败")
		return
	}
	c.String(200, string(png))
}

func (con DefaultController) Qrcode2(c *gin.Context) {
	savePath := "static/upload/qrcode.png"
	err := qrcode.WriteFile("https://www.baidu.com", qrcode.Medium, 256, savePath)

	if err != nil {
		c.String(200, "生成失败")
		return
	}
	file, _ := ioutil.ReadFile(savePath)
	c.String(200, string(file))
}
