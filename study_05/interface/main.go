package main

//interface 接口
import (
	"fmt"
)

type animal interface{
	speak()
}

type cat struct{
	feed int
}

func (c cat)speak()  {
	fmt.Println("喵喵喵~")
}

type dog struct{
	age int 
	name string 
}
func (d dog)speak()  { // 这总是支持指针跟值类型的
	fmt.Println("汪汪汪~")
}

// func (d *dog)speak()  {
// 	fmt.Println("1汪汪汪~")
// }

func main() {
	var a animal
	c1 := cat{
		4,
	}
	fmt.Printf("%T ,%v\n",a,a)
	a = c1
	fmt.Printf("%T ,%v\n",a,a) 
	/*
		这里其实是可以看出 interface 的存储结构是分为值跟值类型的
		而且会更加赋值的类型的替换跟替换值类型，然后再替换值
	*/
	a.speak()
	d1 :=&dog{
		10,
		"ah",
	}
	
	fmt.Printf("%T ,%v\n",d1,d1)
	d1.speak()
	a = d1
	fmt.Printf("%T ,%v\n",a,a)
	a.speak()
	// 值接受者 实现接口 

	var a1 animal
	d2 := dog{
		10,
		"ah2",
	}
	a1 = d2 // 值调用是可以的 
	a1.speak()

	


}