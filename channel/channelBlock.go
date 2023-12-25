package channel

import "fmt"

func ChannelTest() {
	c := make(chan int)
	c <- 42
	fmt.Println("", c)
}
