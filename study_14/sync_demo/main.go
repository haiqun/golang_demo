package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 异步处理请求
func main() {

	r:= gin.Default()

	r.GET("log_sync", func(c *gin.Context) {
		// 异步处理
		// 另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本 [框架要求，不然异步会失效？]
		copyContext := c.Copy()
		go func() {
			time.Sleep(time.Second * 5)
			log.Println("异步执行操作1", copyContext.Request.URL.Path)
		}()
		go func() {
			time.Sleep(time.Second * 3)
			log.Println("异步执行操作2", copyContext.Request.URL.Path)
		}()
		// 同步返回
		c.JSON(200,gin.H{"msg":"成功"})
	})


	r.Run(":9797")

}
