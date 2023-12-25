package channel

import (
	"fmt"
)

func EvenOddChannel() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go sending(even, odd, quit)
	receiving(even, odd, quit)
}

func receiving(even, odd, quit <-chan int) {

	for {
		select {
		case v := <-even:
			fmt.Println("this is even number", v)
		case v := <-even:
			fmt.Println("this is even number", v)
		case v := <-even:
			fmt.Println("this is even number", v)
			return
		}
	}

}

func sending(even, odd, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}

	}
	close(even)
	close(odd)
	q <- 0
}
