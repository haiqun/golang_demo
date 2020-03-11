package main

// 指针
func main() {
	// 1 . & 取地址
	// 2 . * 根据地址取值
	// n := 19
	// fmt.Println(&n)
	// p := &n
	// fmt.Printf("%T\n", p) // *int  int 类型的指针
	// fmt.Println(*p)
	// var a = [...]int{1, 2}
	// fmt.Println(&a) //
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(&a[i]) // 每个数组的元素内存地址都不一致
	// }

	// var a *int     // 声明变量的类型为int的内存地址
	// fmt.Println(a) // 默认是 nil
	// var a1 = new(int)
	// fmt.Println(a1)  // 这里就是一个内存地址 0xc0000a8008
	// fmt.Println(*a1) // 这里就是一个内存地址 0xc0000a8008
	// *a1 = 100        //
	// fmt.Println(*a1) // 这里就是一个内存地址 0xc0000a8008

	// new 和 make 都是用来申请内存的
	// new 很少用 ，一般给基本的数据类型，申请内存空间 如 string ，int 等 ，返回的是对应类型的指针
	// make 用来给slice 。map 、chan 申请内存的 ，make 函数放回的是对应的类型本身

	

}
