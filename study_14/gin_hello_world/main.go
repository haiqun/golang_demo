package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func main() {
	// 创建一个路由 这种方式默认使用了 Logger(),Recovery()中间件
	r := gin.Default()
	// 也可以是用原来的,不带中间件
	//r := gin.New()
	// 绑定路由规则 ，执行的函数
	// 定义一个get
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"hello world ! ")
	})

	//当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{ // 定义返回的状态码，这里是404 ，也可以是 200
			// H 的参数类型是这个 map[string]interface{}
			"message": "Hello world!",
			//"message": &rInfo,
		})
	})
	// 定义api的get 带参数
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK,`hello world ~ `+name)
	})
	// * 号，后续可拼接多参数，但是都放在 action这里
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK,`hello world ~ `+name + action)
	})

	// url 参数 :http://127.0.0.1:9000/welcom?name=lufei
	r.GET("/welcom", func(c *gin.Context) {
		name := c.DefaultQuery("name","lufei")
		action := c.Query("action")
		c.String(http.StatusOK,`hello ~ `+name + " ; action :" + action )
	})

	// 定义一个post
	r.POST("/xxx", func(c *gin.Context) {
		name := c.PostForm("username")
		addr := c.PostForm("address")
		file ,err := c.FormFile("file_name")
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"message" : err.Error(),
			})
		}
		log.Println(file.Filename)
		// 设置存储路径
		dst := fmt.Sprintf("/Users/haiqunfan/work1/tpm/file/%s",file.Filename)
		// 上传到指定的目录
		err = c.SaveUploadedFile(file,dst)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"message" : err.Error(),
			})
		}
		c.JSON(http.StatusOK,gin.H{
			"message" : "ok",
			"username" : name,
			"address" : addr,
		})
	})

	// 多文件上传
	r.POST("/muitipart", func(c1 *gin.Context) {
		// 处理multipart forms提交文件时默认的内存限制是32 MiB
		// 可以通过下面的方式修改
		r.MaxMultipartMemory = 8 << 20  // 设置上传内容的大小为 8 MiB
		form ,_ := c1.MultipartForm()
		// form 获取所有图片
		files :=  form.File["file_name"]  // 接受的文件名
		// 遍历所有图片
		for index,file := range files{
			// 在consolog 输出文件的名字
			log.Println(file.Filename)
			dst := fmt.Sprintf("/Users/haiqunfan/work1/tpm/file/%d_%s", index, file.Filename)
			// 上传文件到指定的目录
			c1.SaveUploadedFile(file, dst)
			//c1.SaveUploadedFile(file, file.Filename) // 默认保存到根目录
		}
		c1.JSON(http.StatusOK,gin.H{
			"message" : "上传成功",
			"code" : 200,
		})
	})



	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	//r.Run()
	// 设置端口
	_ = r.Run(":9000")
}
