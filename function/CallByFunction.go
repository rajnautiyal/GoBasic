package function

func TestMyCallBackFunction() {

	performOperation(13, 12, add)
	performOperation(13, 10, sub)

}

func performOperation(x, y int, f func(int, int) int) int {
	return f(x, y)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
