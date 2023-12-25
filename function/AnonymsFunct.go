package function

import "fmt"

func Foobaar() {

	func() {
		fmt.Println("this is the anonymosu function")
	}()

	func(s string) {
		fmt.Println("I am prinitng the value of s", s)
	}("r")
}
