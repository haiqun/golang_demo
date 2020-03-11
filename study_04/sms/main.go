package main

import (
	"fmt"
	"os"
)

// 学生管理系统

var smr = studentMgr{
	allStudent : make(map[int64]student, 100),
}

func getMenu(){
	fmt.Println(`
1 . 查看全部学员
2 . 新增学员
3 . 删除学员
4 . 退出
	`)
}


func main() {
	for {
		getMenu()
		fmt.Printf("请输入数字，选择你要操作的功能：")
		// 2.1 获取用户的输入值
		var i int
		fmt.Scanln(&i)
		switch i {
			case 1:
				smr.showStudents()
			case 2:
				smr.addStudent()
			case 3:
				smr.deleteStudent()
			case 4:
				os.Exit(200)
			default:
				fmt.Println("没有这个选择项")
		}
	}
}
