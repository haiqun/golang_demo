package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

// 多中数据格式的响应


func main() {

	r := gin.Default()

	// json 格式
	r.GET("/someJson", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"msg" : "someJson hello",
		})
	})
	// 结构体 ，也是按json的方式返回
	r.GET("/someStruct", func(c *gin.Context) {
		var person struct{
			Name string // 结构体要大写，才能被解析
			Age int
		}
		person.Name = "lufei"
		person.Age = 17
		c.JSON(200,person)
	})
	// xml
	r.GET("/someXml", func(c *gin.Context) {
		c.XML(200,gin.H{"xml":"data-xml"})
	})

	// yaml :是个文件格式来的，访问就是下周
	r.GET("/someYaml", func(c *gin.Context) {
		c.YAML(200,gin.H{"yaml":"docker"})
	})

	// protobuf 格式 ，谷歌开发的高效的存储读取工具
	r.GET("/someProtobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})

	r.Run(":9494")

}
