package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入不使用是，调用了包的init()的方法
	"github.com/jmoiron/sqlx"          // 采用这个包
)

// sqlx  的增删改查
// 查询的有了自动映射 但是，需要调整结构体的名字
// 修改，删除，新增的 Exec 函数没有不一样
// 事务的执行，不需要关注执行过程的报错，只关注 开始于提交时的错误判断 ；Beginx 。MustExec

// 创建一个共有的连接池
var db *sqlx.DB // 这里的类型也要调整

// 创建一个user的结构体用来接收数据库操作返回的值
type user struct {
	Id   int // 用sqlx 查询的时候做了映射，这里需要改为大写的首字母
	Name string
	Age  int
}

func initDB() (err error) {
	dsn := "root:fit4hiii@tcp(127.0.0.1:3308)/go_test"
	//链接数据库 - 直接连接不用ping
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	// 设置最大连接数 ：如果拿不到连接，会一直柱塞，等待其他人释放连接
	db.SetMaxOpenConns(10)
	// 设置最大闲置连接数 ：业务空闲的时候可以设置小一点
	db.SetMaxIdleConns(10)
	return nil
}

// sqlx 新增修改的函数没有改变
func transactionDemo() {
	// 开始事务 原来的是 Begin ----- sqlx的是 Beginx
	tx, err := db.Beginx()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Println("开启事务失败:", err)
		return
	}
	// 先新增
	sql := `insert into g_user (name,age) value ("哈哈",8)`
	tx.MustExec(sql)
	/*
		MustExec() 简化了每次执行后的判断，如果失败的 Rollback
	*/
	// 再修改
	sql1 := `update g_user set name = "我是655",age = 101 where id = 10 `
	tx.MustExec(sql1)
	// 提交事务
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("提交事务有误:", err)
		return
	}
	fmt.Println("操作成功！")
}

func getOne(id int) {
	// sqlx 查询
	sql1 := "select id,name,age from g_user where id = ?"
	var u user
	// 不需要按查询字段填写结构体的键值，直接查询，自动映射
	err := db.Get(&u, sql1, id)
	if err != nil {
		fmt.Println("sqlx get err", err)
	}
	fmt.Println(u)
}

func getAll() {
	// sqlx 查询
	sql1 := "select id,name,age from g_user where id > 2"
	var u []user
	// 不需要按查询字段填写结构体的键值，直接查询，自动映射
	err := db.Select(&u, sql1)
	if err != nil {
		fmt.Println("sqlx get err", err)
	}
	fmt.Printf("%#v", u)
}

func main() {
	// 初始化数据库
	err := initDB()
	if err != nil {
		fmt.Printf("初始化db有误！err:%s\n", err)
		return
	}
	// 新增一条数据
	// transactionDemo()
	// 查询
	getOne(1)
	// 查询全量
	getAll()
}
