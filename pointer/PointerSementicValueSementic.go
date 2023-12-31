package pointer

import "fmt"

func TestFun() {

	fmt.Println("this is example of value Sementics")
	v := 1
	fmt.Println(addValueSemantic(v))
	fmt.Println(v)

	fmt.Println("this is example of pointer semetics")
	p := 1
	fmt.Println(addPointerSemantic(&p))

	fmt.Println(p)
}

func addValueSemantic(v int) int {
	v = v + 1
	return v
}

func addPointerSemantic(v *int) int {
	*v = *v + 1
	return *v
}
