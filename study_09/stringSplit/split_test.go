package stringSplit

import (
	"reflect"
	"testing"
)

// 测试组

type testCase struct {
	str  string
	sep  string
	want []string
}

// 测试组

// func TestSplit(t *testing.T) {
// 	testGroup := []testCase{
// 		testCase{"abc", "b", []string{"a", "c"}},
// 		testCase{"a:b:c", ":", []string{"a", "b", "c"}},
// 		testCase{"我是海贼王", "是", []string{"我", "海贼王1"}},
// 	}
// 	for _, v := range testGroup {
// 		got := Split(v.str, v.sep)
// 		// 因为slice不能比较直接，借助反射包中的方法比较
// 		if !reflect.DeepEqual(got, v.want) {
// 			//  测试失败输出错误提示
// 			t.Errorf("got : %s,want :%s ", got, v.want)
// 		}
// 	}

// }

// 子测试
// 外包执行单个子的命令是： go test -v -run=TestSplit/case1
func TestSplit(t *testing.T) {
	testGroup := map[string]testCase{
		"case1": testCase{"abc", "b", []string{"a", "c"}},
		"case2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case3": testCase{"我是海贼王", "是", []string{"我", "海贼王"}},
	}
	for name, v := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(v.str, v.sep)
			// 因为slice不能比较直接，借助反射包中的方法比较
			if !reflect.DeepEqual(got, v.want) {
				//  测试失败输出错误提示
				t.Errorf("got : %s,want :%s ", got, v.want)
			}
		})
	}
}

// -coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件
// go test -cover -coverprofile=c.out
// 将覆盖率相关的信息输出到当前文件夹下面的c.out文件中
// 执行go tool cover -html=c.out，使用cover工具来处理生成的记录信息，
// 该命令会打开本地的浏览器窗口生成一个HTML报告

// 基准测试
// 每个基准测试至少运行1秒
//- 默认跑 go test 不会执行，需要指定-bench参数
// go test -bench=Split
// -benchmem 来获得内存分配的统计数据
// 测试函数以 Benchmark 开头 参数为 *testing.B
// 执行结果解析：
/*
BenchmarkSplit1-4  数字 4 表示 GOMAXPROCS 的值
10000000  199 ns/op 描述调用 10000000次调用的平均值 耗时 199ns
112 B/op 表示每次操作内存分配了112字节
3 allocs/op  每次操作进行了3次内存分配
*/
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

//  性能比较测试
// 多一层封装，用n来控制调用次数，不能直接 用 b.N 这种方式
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

// 为了多跑几次 fib40 的，可以将 benchtime 的执行时间调整到20秒
// go test -bench=Fib40 -benchtime=20s

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }

// func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
