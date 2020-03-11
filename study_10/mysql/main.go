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

//  queryOne 查询一条数据
func queryOne(id int) (u user, err error) {
	// 查询一条数据
	sql1 := "select id,name,age from g_user where id = ?"
	var u1 user
	row := db.QueryRow(sql1, id)
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	// Scan 内部有一个 close 的方法，会释放连接
	// scan 参数的顺序必须跟上面的sql查询的数据顺序一致
	err = row.Scan(&u1.id, &u1.name, &u1.age) // , &u1.name, &u1.age
	if err != nil {
		return u1, err
	}
	return u1, nil
}

func queryAll(id int) {
	sql2 := "select id,age,name from g_user where id > ?"
	rows, _ := db.Query(sql2, id)
	// 手动关闭rows持有的数据库连接
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u2 user
		// scan 参数的顺序必须跟上面的sql查询的数据顺序一致
		err := rows.Scan(&u2.id, &u2.age, &u2.name)
		if err != nil {
			fmt.Println("rows.Scan err", err)
			return
		}
		fmt.Println(u2)
	}
}

func insert(name string, age int) {
	// sql
	sql := `insert into g_user (name,age) value (?,?)`
	ret, err := db.Exec(sql, name, age)
	if err != nil {
		fmt.Printf(" insert into :%s \n", err)
	}
	id, _ := ret.LastInsertId()
	fmt.Println("新增的数据id：", id)
}

func update(name string, id int) {
	sql := "update g_user set name = ? where id = ? "
	ret, err := db.Exec(sql, name, id)
	if err != nil {
		fmt.Println("update :", err)
	}
	// 影响行数
	n, _ := ret.RowsAffected()
	fmt.Println("影响行数：", n)
}

func delete(id int) {
	sql := `delete from g_user where id = ?`
	ret, err := db.Exec(sql, id)
	if err != nil {
		fmt.Println("delete :", err)
	}
	n, _ := ret.RowsAffected()
	fmt.Println("影响行数：", n)
}

// 预处理插入
func prepareInsert() {
	// sql预处理
	sql := `insert into g_user (name,age) value (?,?)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(" Prepare :", err)
	}
	// 关闭连接
	defer stmt.Close()
	// 执行操作
	users := map[string]int{
		"飞":   12,
		"历史上": 22,
		"香吉士": 19,
		"布拉奇": 18,
	}
	for k, v := range users {
		ret, err := stmt.Exec(k, v)
		if err != nil {
			fmt.Println(" Prepare Exec :", err)
		}
		id, _ := ret.LastInsertId()
		fmt.Println("最后执行返回id：", id)
	}

}

func prepareQueryDemo(id int) {
	sqlStr := "select id, name, age from g_user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(" Prepare :", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Println(" Prepare Query:", err)
	}
	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name, &u.age)
		fmt.Println(u)
	}
}

func main() {
	// 初始化数据库
	err := initDB()
	if err != nil {
		fmt.Printf("初始化db有误！err:%s\n", err)
		return
	}
	// 新增一条数据
	insert("test", 20)
	var user1 user
	// 查询一条数据
	user1, err = queryOne(1)
	if err != nil {
		fmt.Printf("queryOne err :%s\n", err)
		return
	}
	fmt.Println(user1)
	// 查询多条数据
	queryAll(1)
	// 修改一条数据
	update("布鲁斯", 4)
	// 删除一条数据
	delete(5)
	// Prepare方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令
	// 预处理新增
	prepareInsert()
	// 查询预处理
	prepareQueryDemo(1)
}
