package main

import (
	"github.com/gin-gonic/gin"
)

// html模板渲染
// http://127.0.0.1:9595/index

func main() {
	r := gin.Default()
	// templates 配置全部的文件夹
	r.LoadHTMLGlob("templates/*")
	// 配置单个文件
	//r.LoadHTMLFiles("templates/index.html")
	r.GET("/index", func(c *gin.Context) {
		// 根据文件名渲染
		// 最终是json将title替换的
		c.HTML(200,"index.html",gin.H{"title":"我的标题"})
	})
	r.Run(":9595")
}
