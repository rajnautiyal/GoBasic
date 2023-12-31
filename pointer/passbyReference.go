package pointer

import "fmt"

func TestByRefrence() {
	x := 42
	delta(&x)
	fmt.Println(x)

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	changeSliceValue(xi)
	fmt.Println(xi)

}
func delta(x *int) {
	*x = 43
}

func changeSliceValue(slice []int) {
	slice[3] = 100
}
