package main

import (
	"fmt"
)

var x int

type person struct{
	first string
	last string
	
}
type secretAgent struct{
	person
	ltk bool
}
func (p person) speak(){
	fmt.Println("I am", p.first, p.last, " _ the person speak")
}
func (s secretAgent ) speak(){
	fmt.Println("I am", s.first, s.last, " _ the secretAgent speak")
}
type human interface{
	speak()
}

func bar(h human){
	switch h.(type){
		case person: fmt.Println("I was passed into barrr ", h.(person).first)
		case secretAgent: fmt.Println("I was passed into barrr ", h.(secretAgent).first)
		
	}
}

func main() {
sa1 := secretAgent{
		person: person{
			 "James",
			 "Bond",
			
			},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			 "Miss",
			 "Moneypenny",
			
			},
		ltk: true,
	}
	
	p1 := person{
	first: "Dr.",
	last: "Yes",}
	bar(sa1)
	bar(sa2)
	bar(p1)

	
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}





func foo2 (s string, x ...int ) int{
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
	fmt.Println(len(x))
	fmt.Println(cap(x))
	
	sum := 0
	
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position, ", i, "we are now adding, ", v, "to the total")
		fmt.Println("sum = ", sum)
	}
	return sum
}


func mouse(fn string, ln string) (string, bool){
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false
	return a,b
}


func foo1(s string) string{

	sa1 := person{
		first: "hello",
		last: "friend",
		
	}
	
	return fmt.Sprint("My info: ",sa1,  s)
}


func bar1(){
	fmt.Println("hello Jane")
}

