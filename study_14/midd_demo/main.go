package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 中间件练习，统计请求所用的时间 

func myTime(c *gin.Context)  {
	t := time.Now()
	// 执行Abort()之后，就会继续执行下面的程序
	//c.Abort()
	c.Next()
	fmt.Println("请求用时：",time.Since(t))
}

func myLog()gin.HandlerFunc  {
	return func(c *gin.Context) {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("time:%s\n",time.Now())
	}
}

func main()  {
	r := gin.Default()
	// 使用全局中间件
	r.Use(myTime)
	// 定义一个组路由
	testGroup := r.Group("/test")
	{
		testGroup.GET("/index",indexHandler)
		testGroup.GET("/home",homeHandler)
	}
	// 这样也命中
	r.GET("/home/info", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		c.JSON(200,gin.H{
			"msg" : "test~",
		})
	})

	// 有局部的先执行局部的，再执行全局的中间件
	r.GET("/index/log",myLog(), func(c *gin.Context) {
		c.JSON(200,gin.H{
			"msg" : "log~",
		})
	})

	r.Run(":9999")
}

func indexHandler(c *gin.Context)  {
	time.Sleep(3 * time.Second)
	c.JSON(200,gin.H{
		"msg" : "hello~",
	})
}

func homeHandler(c *gin.Context)  {
	time.Sleep(50*time.Millisecond)
	c.JSON(200,gin.H{
		"msg" : "world~",
	})
}

