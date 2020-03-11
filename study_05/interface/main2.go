package main

//多个类型实现同一接口
import "fmt"

type mover interface{
	move()
}

type cat struct{
	name string
}

func (c cat)move()  {
	fmt.Printf("奔奔跳跳的%s \n",c.name)
}

type car struct{
	brand string
}

func (c car)move()  {
	fmt.Printf("速度 70 迈 --- %s \n",c.brand)
}

type WashingMachine interface {
	dry() 
	wash()
}

type dryer struct{

}

type haier struct{
	slogan string
	dryer
}

func (d dryer)dry()  {
	fmt.Println("甩甩甩-全部甩干水")
}

func (h haier)wash()  {
	fmt.Println(h.slogan)
}

func main() {
	// test1
	// var m mover
	// cat1 := cat{
	// 	name:"阿喵",
	// }
	// m = cat1
	// m.move()
	// car1 := car{
	// 	brand :"宝马",
	// }
	// m = car1
	// m.move()

	// test2
	var w1 WashingMachine
	// dry1 := dryer{}
	haier1 := haier{
		slogan : "洗一洗，更将康",
	}
	w1 = haier1
	w1.dry()
	w1.wash()
}