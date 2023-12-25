package channel

import "fmt"

func RangeChannel() {
	c := make(chan int)
	go putValue(c)
	receiveValue(c)
}

func putValue(c chan<- int) {

	for i := 0; i < 100; i++ {
		c <- 42
	}
	close(c)
}

func receiveValue(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
