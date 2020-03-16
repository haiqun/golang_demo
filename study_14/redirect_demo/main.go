package main

import (
	"github.com/gin-gonic/gin"
)

// 重定向

func main() {
	r := gin.Default()

	r.GET("/redirect", func(c *gin.Context) {
		// 要写 http ，不然属于内容的跳转 ，
		// 内外部的调转用的是同一个方法
		c.Redirect(302,"http://www.baidu.com")
	})

	r.Run(":9696")
}
