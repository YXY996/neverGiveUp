package middleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleWare(c *gin.Context) {
	//middleWare 中间件 在 访问handler方法前进行处理
	//c.Set("username", "yxy")
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)

	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done in path", cCp.Request.URL)
	}()
}
