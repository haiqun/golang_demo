package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	// 创建一个路由
	r:=gin.Default()
	// 定义一个get
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg" : "hello world~",
			"code" : 0,
		})
	})

	r.Run(":8080")
}
