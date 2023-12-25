package function

import "fmt"

func TestRecursion() {
	x := factorial(3)
	fmt.Printf("this is the value %d", x)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
