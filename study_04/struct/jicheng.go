package main

import "fmt"

// 机构体模拟实训继承

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s ,会动\n", a.name)
}

type dog struct {
	age int
	animal
}

func (d dog) wang() {
	fmt.Printf("%s 会叫， 汪汪汪~", d.name)
}

func main() {
	p1 := dog{
		age:    4,
		animal: animal{name: "lufei"},
	}
	// fmt.Println(p1)
	p1.move()
	p1.wang()
}
