package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
)

// 普通请求
func f1() {
	resp, err := http.Get("http://127.0.0.1:9090/info/?name=lufei&age=19")
	if err != nil {
		fmt.Printf("请求有误err：%s\n", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应信息有误err：%s\n", err)
		return
	}
	fmt.Println(string(b))
}

// 定制请求
func f2() {
	// url_encode
	// 获取url对象
	date := url.Values{}
	urlObj, err := url.Parse("http://127.0.0.1:9090/info/")
	if err != nil {
		fmt.Printf("获取url对象有误 %s\n", err)
		return
	}
	// 构建请求数据
	date.Set("name", "路飞")
	date.Set("age", "17")
	queryStr := date.Encode()
	fmt.Println(queryStr) // age=17&name=%E8%B7%AF%E9%A3%9E
	// 合并完成请求url
	urlObj.RawQuery = queryStr
	fmt.Println(urlObj) // http://127.0.0.1:9090/info/?age=17&name=%E8%B7%AF%E9%A3%9E
	// 生成请求
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	// fmt.Println(req)
	// 发起请求
	// resp, err := http.DefaultClient.Do(req)
	// 发起一个短连接 （禁用长链接）
	// 如果不这样设置，那么存储会存在很多链接 ，占用链接数
	// 也可以抽离为公共变量
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发起链接有误：err", err)
	}

	defer resp.Body.Close() // 释放连接
	//打印响应的内容
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取用户传过来的信息有误：err", err)
	}
	// fmt.Println(ioutil.ReadAll(r.Body)) //请求的数据
	fmt.Println(string(data)) //请求的数据
}

func main() {
	// f1() // 普通的get请求
	f2()
}
