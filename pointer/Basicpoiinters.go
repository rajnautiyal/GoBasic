package pointer

import "fmt"

func Pointer() {
	s := "james"
	fmt.Printf("this is the type for %v the point %T \n", &s, &s)

	y := &s

	fmt.Printf("this is the address of x %v \n", y)

	fmt.Printf("this is the value of y %v \n", *y)

	fmt.Printf("this is the value of s %v \n", *&s)

	*y = "Inesh"
	fmt.Printf("this is the value of s %v \n", s)

}
