package main

import (
	"fmt"
)

//  for range : 如果采用for range ，会对原slice的值进行复制，这样改变复制之后的值，对原slice是不会产生影响的
func f1()  {
	b := []int{1, 2}
	// 场景一  这个例子是为了测试for range. 说明for range 每次会将slice的内容copy给v, 所以v变化，也不会影响b
	fmt.Println("场景一 无影响")
	for _, v := range b {
		v++
		//fmt.Println(k, v)
	}
	for k, v := range b {
		fmt.Println(k, v)
	}
	fmt.Println("=============================")
}

//  for range
func f2()  {
	b := []int{1, 2}
	// 场景二 要想产生实际影响，要这样做
	fmt.Println("场景二 产生影响")
	for k, v := range b {
		b[k] = v + 1
	}
	for k, v := range b {
		fmt.Println(k, v)
	}
	fmt.Println("=============================")
}

// copy : copy的作用是将源slice copy给目的slice，类型必须一致
func f3()  {
	fmt.Println("slicecopy 方法最终的复制结果取决于较短的那个切片，当较短的切片复制完成，整个复制过程就全部完成了。")
	array := []int{10, 20, 30, 40}
	slice := make([]int, 6)
	n := copy(slice, array)
	fmt.Println(n,slice)
	fmt.Println("=============================================")

	fmt.Println("即使array2有7个元素，但是slice2只接收了6个。")
	array2 := []int{10, 20, 30, 40, 50,60,70}
	slice2 := make([]int, 6)
	n2 := copy(slice2, array2)
	fmt.Println(n2,slice2)
	fmt.Println("=============================================")
}

// slice作为参数
func f4()  {
	arrayA := [2]int{100, 200}
	sliceA := arrayA[:]
	fmt.Printf("sliceA : %p , %v\n", &sliceA, sliceA)
	f4a(sliceA)
	fmt.Println("---------------------------")
}

// 如果直接将slice 作为参数传进来，会有一个值拷贝的过程，如果这个slice 过大，会影响性能
func f4a(x []int) {
	fmt.Printf("func Slice : %p , %v\n", &x, x)
}

// append 操作 = append之后的slice跟之前的slice肯定不是一个
func f5()  {
	arr := [3]int{1,2,3}
	slice := arr[0:2] // 左包右不包
	newSlice := append(slice, 50)
	fmt.Printf("arrPoint = %p, slicePointer = %p, newSlicePoint = %p \n", &arr, &slice, &newSlice)
	// 利用array新生成的slice, 以及利用slice append操作生成的newSlice 地址都是不一样的
	fmt.Println("=============================================")
}

// append之后，底层数组变了吗？ -- 如果没有操作容量，不变
func f6()  {
	arr := [3]int{1,2,3}
	slice := arr[0:2]
	newSlice := append(slice, 10)
	fmt.Printf("arrPoint = %p, slicePointer = %p, newSlicePoint = %p \n", &arr, &slice, &newSlice)
	fmt.Printf("arrPoint = %p, slice array Pointer = %p, newSlice array Point = %p \n", &arr[0], &slice[0], &newSlice[0])
	fmt.Println("=============================================")
}

func main() {
	//f1()
	//f2()
	//f3()
	//f4()
	//f5()
	f6()
}
