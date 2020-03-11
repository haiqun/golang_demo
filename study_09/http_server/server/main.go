package main

import "net/http"

import "io/ioutil"

import "fmt"

// net/http server

// 设置跳转的方法 ，参数 func(ResponseWriter, *Request)
func hhh(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./hello.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
	}
	// 参数是 []byte => 强制装换的语法是 type() , 因为string底层就是切片，所以可以直接转换
	w.Write(b)
}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)    //请求的url
	fmt.Println(r.Method) //请求的方法
	// ioutil.ReadAll 参数 r io.Reader
	// 是一个接口类型，接口条件是 ：Read(p []byte) (n int, err error)
	query := r.URL.Query()
	fmt.Println(query)             //请求的url的参数 // map[age:[19] name:[lufei]]
	fmt.Println(query.Get("name")) //请求的url的参数 // lufei
	fmt.Println(r.URL.Path)        //请求的url的uri
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("读取用户传过来的信息有误：err", err)
	}
	// fmt.Println(ioutil.ReadAll(r.Body)) //请求的数据
	fmt.Println(string(data)) //请求的数据
	w.Write([]byte("ok"))
}

func main() {
	// 创建一个请求链接
	url1 := "/test/goto/hhh/"
	http.HandleFunc(url1, hhh)
	url2 := "/info/"
	http.HandleFunc(url2, info)
	// 创建一个http的链接
	http.ListenAndServe("127.0.0.1:9090", nil) // 这里的nil参数类型是一个interface
}
