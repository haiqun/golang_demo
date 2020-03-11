package main

import (
	"fmt"
	"time"
)

// time的场景使用
func main() {
	// 获取8小时之后
	now := time.Now()
	fmt.Println(now)

	hh, _ := time.ParseDuration("8h")
	hh1 := now.Add(hh)
	fmt.Println(hh1)
	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1)
	// 30秒之后
	s, _ := time.ParseDuration("30s")
	s1 := now.Add(s)
	fmt.Println(s1)
	// 5分支之前
	m, _ := time.ParseDuration("-5m")
	m1 := now.Add(m)
	fmt.Println(m1)

	// Sub 计算两个时间差 并按 分钟的方式返回
	subM := now.Sub(d1)
	fmt.Println(subM.Minutes(), "分钟")
	// Sub 计算两个时间差 并按 秒钟的方式
	subM1 := now.Sub(d1)
	fmt.Println(subM1.Seconds(), "秒钟")
	fmt.Println(subM1.Seconds())

	// 倒转过来减
	subM2 := m1.Sub(now)
	fmt.Println(subM2.Seconds())
	// 相减的时间格式都是这种的：2020-03-02 22:19:12.048317 +0800 CST m=-86399.999736444
	subM3 := m1.Sub(d1) // 五分之前的时间 减去 1天前的时间
	fmt.Println(subM3.Seconds())

	// 把一个字符串的时间格式，变成指定时区的时间,格式，并计算相差的时间
	// 获取指定时区的时间如东八区
	loc, err := time.LoadLocation("Asia/shanghai")
	if err != nil {
		fmt.Printf("获取指定时区的指针有问题 : %s\n", err)
	}
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2050-03-03 00:00:00", loc)
	if err != nil {
		fmt.Printf("字符串按指定格式，指定时区转换的时间，装换报错 : %s\n", err)
	}
	subM4 := timeObj.Sub(now)
	fmt.Println(subM4.Hours() / 24 / 365) // 算出距离现在的时间 hours day 
}
