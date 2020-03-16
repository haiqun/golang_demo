package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 一个可以匹配所有请求方法的Any方法
func main() {
	r := gin.Default()
	// 也可以是用原来的,不带中间件
	//r := gin.New()
	// 绑定路由规则 ，执行的函数
	// 定义一个get
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"hello world ! ")
	})

	r.Any("/test", func(c *gin.Context) {
		c.String(http.StatusOK,"hello test ! ")
	})

	r.Run(":9292")
}
