package main
// 接口嵌套
import "fmt"
// type sayer interface{
// 	say()
// }
// type mover interface{
// 	move()
// }

// type animal1 interface{
// 	mover
// 	sayer
// }

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

type dog struct {
	name string
}

func (d dog)move()  {
	fmt.Printf("%s , 奔奔跳跳的~ \n",d.name)
}

func (d dog)say()  {
	fmt.Printf("%s 说, 我会唱跳rap~ \n",d.name)
}

func main()  {
	var a animal
	// fmt.Println(a)
	d1 := dog{
		"ah",
	}
	a = d1
	a.move()
	a.say()
}