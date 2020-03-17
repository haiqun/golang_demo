package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入不使用是，调用了包的init()的方法
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB() (err error)  {
	dsn := "root:fit4hiii@tcp(127.0.0.1:3308)/go_test"
	db ,err = sqlx.Connect("mysql",dsn)
	if err != nil {
		return err
	}
	// 设置最大连接数
	db.SetMaxOpenConns(100)
	// 设置最大空闲数
	db.SetMaxIdleConns(10)
	return
}

func queryAllBook() (list []*Book, err error) {
	sqlStr := `select id,title,price from book`
	err = db.Select(&list, sqlStr)
	if err != nil {
		return
	}
	return
}

func insertBook(title string,price string)(err error)  {
	sqlStr := `insert into book(title,price) value(?,?)`
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("录入失败")
		return
	}
	return
}

func delBook(id int) (err error) {
	sqlStr := `delete from book where id = ?`
	fmt.Println("id等于：",id)
	_, err = db.Exec(sqlStr,id)
	if err != nil {
		fmt.Println("删除失败")
		return
	}
	return
}