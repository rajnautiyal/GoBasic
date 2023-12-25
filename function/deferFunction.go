package function

import "fmt"

func Flowdefer() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("this the foo method")
}

func bar() {
	fmt.Println("this the bar method")
}
