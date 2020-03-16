package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 路由组 v1
	v1 := r.Group("/v1")
	{
		// 请求这个login 路径必须拼接上组 v1
		v1.GET("/login",login)
	}
	v2 := r.Group("/v2") //{}不能放这里
	{
		v2.POST("/login",loginPost)
	}
	r.Run(":9393")
}

func login(c *gin.Context)  {
	c.String(http.StatusOK,`hello world ~ `)
}

func loginPost(c *gin.Context)  {
	c.String(http.StatusAccepted,"post success ")
}