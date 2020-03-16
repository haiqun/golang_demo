package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
为了能够更方便的获取请求相关参数，提高开发效率，
我们可以基于请求的Content-Type识别请求数据类型并利用反射机制自动提取请求中QueryString、form表单、JSON、XML等参数到结构体中
*/
// Binding from JSON
type Login struct {
	// 限制了form表达接收是 user的字段名叫 username1 ，以json过来时字段叫 user
	User     string `form:"username1" json:"user" binding:"required"` // 设置了必填 required
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	/*
	curl -X POST \
	  http://127.0.0.1:9191/loginJSON \
	  -H 'content-type: application/json' \
	  -d '{"user": "q1mi", "password": "123456"}'
	*/
	// 绑定JSON的示例 ({"user": "root", "password": "123456"})
	router.POST("/loginJSON", func(c *gin.Context) {
		fmt.Println("json")
		// 声明一个结构体
		var login Login
		// 将结构体指针 传入，进行双向绑定
		// 将request中的body数据，按照json格式解析到结构体
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			// 返回json
			// gin.H 封装好的生成 json 数据的工具
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定form表单示例 (username1=root&password=123456)
	router.POST("/loginForm", func(c *gin.Context) {
		// 方式二
		//fmt.Println("form")
		//var login Login
		//// ShouldBind()会根据请求的Content-Type自行选择绑定器
		//if err := c.ShouldBind(&login); err == nil {
		//	fmt.Println(login)
		//	c.JSON(http.StatusOK, gin.H{
		//		"user":     login.User,
		//		"password": login.Password,
		//	})
		//} else {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//}
		// 方式一
		var form Login
		// Bind 默认解析并绑定 form格式的数据
		// 是根据请求头的content—type自动推断的
		if err := c.Bind(&form); err != nil{
			c.JSON(200,gin.H{"status":404})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user":     form.User,
			"password": form.Password,
		})

	})

	// 绑定QueryString示例 (http://127.0.0.1:9191/loginForm?username1=root&password=123)
	router.GET("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.Run(":9191")
}