package main

import "fmt"

type student struct {
	name string
	id  int64
}

type studentMgr struct{
	allStudent map[int64]student
}

func (s studentMgr)newStudent(id int64, name string) student {
	return student{
		id:   id,
		name: name,
	}
}

func (s studentMgr)addStudent(){
	var (
		id   int64
		name string
	)
	fmt.Print("请输入你要添加的学生学号:")
	fmt.Scanln(&id)
	// 判断id是否重复
	ret1 := s.allStudent[id].id
	if ret1 != 0 {
		fmt.Println("你输入的id已经存在")
		s.addStudent()
	}
	fmt.Print("请输入你要添加的学生姓名:")
	fmt.Scanln(&name)
	ret := s.newStudent(id, name)
	s.allStudent[id] = ret
}


func (s studentMgr)editStudent(){
	fmt.Println("编辑")
}

func (s studentMgr)deleteStudent(){
	var (
		id       int64
		idconfem int64
	)
	fmt.Print("请输入你要删除的学生学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入你确认你要删除的学生学号:")
	fmt.Scanln(&idconfem)
	if id == idconfem {
		delete(s.allStudent, id)
	} else {
		var n int
		fmt.Println("两次输入id不一致")
		fmt.Print(
`1 . 继续删除
2 . 退出
请选择操作:`)
		fmt.Scanln(&n)
		if n == 1 {
			s.deleteStudent()
		} else {

		}
	}
}

func (s studentMgr)showStudents(){
	for _, v := range s.allStudent {
		fmt.Printf("学生学号：%d 学生姓名：%v\n", v.id, v.name)
	}
}