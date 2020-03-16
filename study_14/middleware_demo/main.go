package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 中间件
func middleware()gin.HandlerFunc  {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Println("中间件开始")
		// 设置这次请求的值，可以被接下来的请求获取到
		c.Set("q1","test_lufei")
		c.Next()
		fmt.Println("next之后",time.Now().UnixNano())
		t2 := time.Since(start)
		// 获取q1的值，是被web的目标修改过的
		q1 ,_ := c.Get("q1")
		fmt.Println("中间件结束：",t2,q1)
	}
}

func main() {
	r := gin.Default()
	// 注册中间件
	r.Use(middleware())
	{
		/*
			请求的顺序是 执行中间件 -> 执行next -> web请求 -> next之后的执行
		*/
		r.GET("/index", func(c *gin.Context) {
			fmt.Println("index收到请求了",time.Now().UnixNano())
			// 获取中间件设置的值
			q1 ,_ := c.Get("q1")
			c.JSON(200,gin.H{
				"msg" : "test",
				"data" : q1,
			})
			// 再次修改q1的值
			c.Set("q1","hhh")
		})

		r.GET("/index1", func(c *gin.Context) {
			fmt.Println("index收到请求了",time.Now().UnixNano())
			// 获取中间件设置的值
			q1 ,_ := c.Get("q1")
			c.JSON(200,gin.H{
				"msg" : "test",
				"data" : q1,
			})
			// 再次修改q1的值
			c.Set("q1","hhh")
		})

	}
	r.Run(":9898")
}




