package main

type test1 interface {
	save()
}

type person1 struct {
	name string
}

type person2 struct {
	age int
}

func (p *person1)save()  {

}

func main() {

}