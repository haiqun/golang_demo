package stringSplit

import "testing"

import "reflect"

// 单元测试

// 测试文件在当前目录下 文件名是 (xx)_test.go 结尾 如果没有：[no test files]
// 参数名 必须是 Test 开头
// 参数必须是 t *testing.T
// got 返回值
// want 为预想值

func TestB3Split(t *testing.T) {
	got := Split("abc", "b")
	want := []string{"a", "c1"}
	// 因为slice不能比较直接，借助反射包中的方法比较
	if !reflect.DeepEqual(got, want) {
		//  测试失败输出错误提示
		t.Errorf("got : %s,want :%s ", got, want)
	}
}

func TestB2Split(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	// 因为slice不能比较直接，借助反射包中的方法比较
	if !reflect.DeepEqual(got, want) {
		//  测试失败输出错误提示
		t.Errorf("got : %s,want :%s ", got, want)
	}
}

// 执行命令  go test -cover  测试覆盖了，在测试中至少被运行一次的代码占总代码的比例
// 执行命令  go test  跑动测试脚本
// 执行命令  go test -v  跑动测试脚本的详细信息
// 执行命令  go test -v -run="More"  跑动测试脚本的详细信息
// -run参数，它对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执行

func TestB1Split(t *testing.T) {
	got := Split("a:ksdjfku", "dj")
	want := []string{"a:ks", "fku"}
	// 因为slice不能比较直接，借助反射包中的方法比较
	if !reflect.DeepEqual(got, want) {
		//  测试失败输出错误提示
		t.Fatalf("got : %s,want :%s ", got, want)
	}
}
