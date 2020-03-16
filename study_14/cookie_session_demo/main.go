package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// cookie 与 session
func main() {
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		// 设置cookie
		ret  , err := c.Cookie("lufei")
		if err != nil {
			ret = "null"
			// 获取不到cookie
			v := `你的cookie值是:hhhh _` + time.Now().String()
			// 设置新的 localhost
			c.SetCookie("lufei",v,60,"/","127.0.0.1",false,true)
		}
		c.JSON(200,gin.H{
			"msg" : fmt.Sprintf("你的cookie值是%s",ret),
		})
	})

	r.Run(":8081")
}
