package main

var ch1 chan int

func main()  {
	ch1 = make(chan int,1000)
	for v:= range ch1 {
		println(v)
	}
}