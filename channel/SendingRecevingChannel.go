package channel

import "fmt"

func SendingRecevingChannel() {
	c := make(chan int)
	go send(c)
	go recieve(c)
}

func send(c chan<- int) {
	c <- 42
}

func recieve(c <-chan int) {
	fmt.Println(<-c)
}
