package main

func main() {
	ch1 := make(chan int ,10)
	close(ch1)
	select {
	case  ch1<-1:
		if ok {

		}
	}

}
