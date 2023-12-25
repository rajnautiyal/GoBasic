package channel

import (
	"fmt"
	"sync"
)

func FainIn() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)
	go evenOddValue(even, odd)
	go fanInOneChannel(even, odd, fanIn)

	for v := range fanIn {
		fmt.Println("this is the value of fanIN", v)
	}
}

func evenOddValue(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func fanInOneChannel(even, odd <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range even {
			fanIn <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanIn <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanIn)
}
