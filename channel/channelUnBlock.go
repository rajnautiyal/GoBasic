package channel

import "fmt"

func ChannelUblockTest() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}
