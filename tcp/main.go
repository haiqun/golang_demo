package main

import (
	"fmt"
)

type a interface {
	show()
}

type p struct {
	name string
}

func (p1 * p)show()  {
	fmt.Println("show name : "+p1.name)
}

var _ a = new(p)

func f2()  {
	err := recover()
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("333")
	}
}

func f1()  {
	defer f2()
	panic("test7777")
	defer fmt.Println("123")
}

func main()  {
	//p1 := p{
	//	name:"test",
	//}
	//p1.show()

	//q := make(map[int]string,1)
	////q[1] = "test"
	//fmt.Println(q)

	//a := make([]int,5,10)
	//fmt.Println(cap(a),len(a))
	//for i:=0;i<10;i++  {
	//	a = append(a, i)
	//}
	//
	//for i:=0;i<10;i++  {
	//	a[i] = i
	//}
	//fmt.Println(a,cap(a),len(a))


	f1()

}