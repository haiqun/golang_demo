package main

import (
	"fmt"
	"os"
)

// 学生管理系统
type student struct {
	id   int64
	name string
}

var studentVar map[int64]*student

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showStudentAll() {
	for _, v := range studentVar {
		fmt.Printf("学生学号：%d 学生姓名：%v\n", v.id, v.name)
	}
}

func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入你要添加的学生学号:")
	fmt.Scanln(&id)
	// 判断id是否重复
	ret1 := studentVar[id]
	if ret1 != nil {
		fmt.Println("你输入的id已经存在")
		addStudent()
	}
	fmt.Print("请输入你要添加的学生姓名:")
	fmt.Scanln(&name)
	ret := newStudent(id, name)
	studentVar[id] = ret
}

func deleteStudent() {
	var (
		id       int64
		idconfem int64
	)
	fmt.Print("请输入你要删除的学生学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入你确认你要删除的学生学号:")
	fmt.Scanln(&idconfem)
	if id == idconfem {
		delete(studentVar, id)
	} else {
		var n int
		fmt.Println("两次输入id不一致")
		fmt.Print(
`1 . 继续删除
2 . 退出
请选择操作:`)
		fmt.Scanln(&n)
		if n == 1 {
			deleteStudent()
		} else {

		}

	}
}


func main() {
	studentVar = make(map[int64]*student, 100)
	for {
		// studentVar := make(map[int]student, 100)
		// 1 欢迎用语
		fmt.Println("欢迎光临学生管理系统！")
		// 2 选择功能
		fmt.Println(`
1 . 查看全部学员
2 . 新增学员
3 . 删除学员
4 . 退出
		`)
		fmt.Printf("请输入数字，选择你要操作的功能：")
		// 2.1 获取用户的输入值
		var i int
		fmt.Scanln(&i)
		switch i {
		case 1:
			showStudentAll()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(200)
		default:
			fmt.Println("没有这个选择项")
		}
	}
}
