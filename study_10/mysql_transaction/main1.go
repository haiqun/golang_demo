package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入不使用是，调用了包的init()的方法
)

// mysql 的增删改查

// 创建一个共有的连接池
var db *sql.DB

// 创建一个user的结构体用来接收数据库操作返回的值
type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	dsn := "root:fit4hiii@tcp(127.0.0.1:3308)/go_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	// 设置最大连接数 ：如果拿不到连接，会一直柱塞，等待其他人释放连接
	db.SetMaxOpenConns(10)
	// 设置最大闲置连接数 ：业务空闲的时候可以设置小一点
	db.SetMaxIdleConns(10)
	return nil
}

func transactionDemo() {
	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Println("开启事务失败:", err)
		return
	}
	// 先新增
	sql := `insert into g_user (name,age) value ("哈哈",8)`
	_, err = tx.Exec(sql)
	if err != nil {
		tx.Rollback()
		fmt.Println("新增有误:", err)
		return
	}
	// 再修改
	sql1 := `update g_user set name = "我是655",age = 101 where id = 10 `
	_, err = tx.Exec(sql1)
	if err != nil {
		tx.Rollback()
		fmt.Println("更新操作有误:", err)
		return
	}
	// 提交事务
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("提交事务有误:", err)
		return
	}
	fmt.Println("操作成功！")
}

func main() {
	// 初始化数据库
	err := initDB()
	if err != nil {
		fmt.Printf("初始化db有误！err:%s\n", err)
		return
	}
	// 新增一条数据
	transactionDemo()
}
