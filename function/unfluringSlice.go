package function

import "fmt"

func unfluringSlice() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := Sum(x...)
	fmt.Println("this is the sum", sum)
}
