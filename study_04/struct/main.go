package main

import "fmt"

//结构体 => 自定义类型
// 类型名 ： 标识自定义的结构体的名称，在同一个包内不能重复
// 字段名 ： 表示结构体字段名，结构体中的字段名必须唯一
// 字段类型 ： 表示结构体字段的具体类型
type config struct { // main.config
	host     string
	port     int
	username string // 这里的类型为go的原来的类型 如 string就是string
	password string
	address  []string
}

func main() {
	var db config // 实例化 使用之前一定要实例化
	db.host = "127.0.0.1"
	db.port = 3308
	db.username = "root"
	db.password = "admin-rooo"
	db.address = []string{"192,168,1,1", "192,168,1,2"}
	fmt.Println(db)
	fmt.Println(db.host)
	fmt.Printf("%T\n", db)
	fmt.Printf("%T\n", db.host)
	fmt.Printf("%T\n", db.address)
	var redis config
	redis.host = "192.168.1.1"
	redis.address = []string{"192,168,1,3", "192,168,1,4"}
	fmt.Println(redis)
	// 匿名结构体 : 多用于临时场景
	var s struct {
		x string
		y int
		z []string
	}
	s.x = "i'm lufei"
	s.y = 100
	s.z = []string{"hhhh", "中", "d回到家"}
	for _, y := range s.z {
		println(y)
	}
	fmt.Println(s.z)
	fmt.Printf("%T\n", s)
}
