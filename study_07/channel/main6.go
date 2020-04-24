package main

import (
	"fmt"
	"unsafe"
)

/*
	内存对齐
	不管结构体的字节长度是超过还是不足8个字节，他都会自动填充约束至8个字节
*/

type Example struct {
	BoolValue bool
	IntValue int16
	FloatValue float32
	KKK int64
}

func main() {
	example := &Example{
		BoolValue:  true,
		IntValue:   10,
		FloatValue: 3.141592,
		KKK:10,
	}

	exampleNext := &Example{
		BoolValue:  true,
		IntValue:   10,
		FloatValue: 3.141592,
		KKK:10,
	}

	alignmentBoundary := unsafe.Alignof(example)

	sizeBool := unsafe.Sizeof(example.BoolValue)
	offsetBool := unsafe.Offsetof(example.BoolValue)

	sizeInt := unsafe.Sizeof(example.IntValue)
	offsetInt := unsafe.Offsetof(example.IntValue)

	sizeFloat := unsafe.Sizeof(example.FloatValue)
	offsetFloat := unsafe.Offsetof(example.FloatValue)

	sizeInt64 := unsafe.Sizeof(example.KKK)
	offsetInt64 := unsafe.Offsetof(example.KKK)

	sizeBoolNext := unsafe.Sizeof(exampleNext.BoolValue)
	offsetBoolNext := unsafe.Offsetof(exampleNext.BoolValue)

	fmt.Printf("Alignment Boundary: %d\n", alignmentBoundary)

	fmt.Printf("BoolValue = Size: %d Offset: %d Addr: %v\n",
		sizeBool, offsetBool, &example.BoolValue)

	fmt.Printf("IntValue = Size: %d Offset: %d Addr: %v\n",
		sizeInt, offsetInt, &example.IntValue)

	fmt.Printf("FloatValue = Size: %d Offset: %d Addr: %v\n",
		sizeFloat, offsetFloat, &example.FloatValue)

	fmt.Printf("sizeInt64 = Size: %d Offset: %d Addr: %v\n",
		sizeInt64, offsetInt64, &example.KKK)

	fmt.Printf("Next = Size: %d Offset: %d Addr: %v\n",
		sizeBoolNext, offsetBoolNext, &exampleNext.BoolValue)
}