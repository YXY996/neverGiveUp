package models

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

func MakeCaptcha() (string, string, error) {
	var driver base64Captcha.Driver
	//验证码格式设置
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          2,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driver = driverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	return id, b64s, err
}

func VerifyCaptcha(id string, verifyValue string) bool {
	fmt.Println(id, verifyValue)
	if store.Verify(id, verifyValue, true) {
		return true
	}
	return false
}
