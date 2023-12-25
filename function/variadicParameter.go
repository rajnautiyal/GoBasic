package function

func Sum(numbers ...int) int {
	var sum int = 0
	for _, num := range numbers {
		sum = sum + num
	}
	return sum
}
