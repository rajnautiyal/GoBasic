package channel

import "fmt"

func Chan22() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	fmt.Println("this the the issue", <-c)
	close(c)
}
