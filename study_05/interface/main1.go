package main

// 类型与接口的关系
// 一个类型实现多个接口

import "fmt"

type sayer interface{
	say()
}

type mover interface{
	move()
}

type dog struct{

}

func (d dog)move()  {
	fmt.Println("跑来跑去")
}

func (d dog)say()  {
	fmt.Println("唱，跳，rap")
}

func main() {
	var (
		s sayer
		m mover
	)
	d1 := dog{}
	s = d1
	m = d1
	s.say()
	m.move()
	fmt.Printf("%T %v\n", s,s)
	fmt.Printf("%T %v\n", m,m)
}