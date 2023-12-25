package function

import "fmt"

type Address struct {
	Name     string
	location string
}

type SecretAgent struct {
	address Address
	itk     bool
}

// Interface

type human interface {
	speak()
}

func SaySomething(h human) {
	h.speak()
}
func TestMyInterface() {

	p1 := Address{
		Name:     "rudraprayag",
		location: "ddun",
	}

	s1 := SecretAgent{
		address: Address{
			Name:     "secret agent ",
			location: "ddun",
		},
		itk: true,
	}

	SaySomething(p1)
	SaySomething(s1)

}

func (p SecretAgent) speak() {
	fmt.Println("this is want it is ", p.address.Name)
}

func (p Address) speak() {
	fmt.Println("this is want it is ", p.Name)
}
