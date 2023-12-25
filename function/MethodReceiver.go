package function

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func MethodReceive() {

	p1 := person{
		firstName: "raj",
		lastName:  "nauti",
	}
	p1.speak()
}

func (p person) speak() {
	fmt.Println("this is want it is ", p.firstName)
}
