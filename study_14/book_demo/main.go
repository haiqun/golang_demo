package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main()  {
	r := gin.Default()

	// 初始化数据库
	err := InitDB()
	if err != nil {
		panic("数据库连接失败 :"+err.Error())
	}

	fmt.Println(db)
	// 加载模板
	r.LoadHTMLGlob("./tmp/*")
	r.GET("/", index)
	r.POST("/book/new", add)
	r.GET("/book/new", new)
	r.GET("/book/delete", del)
	_ = r.Run(":8080")
}

func index(c *gin.Context)  {
	// 获取数据
	list ,err  := queryAllBook()
	if err != nil {
		c.JSON(200,gin.H{
			"msg" : err,
			"code" : -1,
		})
	}
	fmt.Println(list)
	// 渲染数据
	c.HTML(200,"list_book.html",gin.H{
		"code" : 0,
		"data" :list,
	})
}

func new(c *gin.Context)  {
	// 渲染数据
	c.HTML(200,"new_book.tmp",gin.H{
		"code" : 0,
	})
}

func add(c *gin.Context)  {
	// 新增数据
	var book Book
	if err := c.Bind(&book); err != nil{
		c.JSON(200,gin.H{"status":404})
		return
	}
	err := insertBook(book.Title, book.Price)
	if err != nil {
		fmt.Println("")
	}
	//c.JSON(200,gin.H{
	//	"code":0,
	//	"msg":"提交成功",
	//})
	c.Redirect(http.StatusMovedPermanently, "/")

}

func del(c *gin.Context)  {
	id := c.DefaultQuery("id","0")
	bid, _ := strconv.Atoi(id)
	err := delBook(bid)
	if err != nil {
		c.JSON(200,gin.H{
			"msg" : err,
			"code" : -1,
		})
	}
	//c.JSON(200,gin.H{
	//	"code":0,
	//	"msg":"删除成功",
	//})
	//time.Sleep(time.Second * 3)
	// 调整回首页
	c.Redirect(http.StatusMovedPermanently, "/")
}