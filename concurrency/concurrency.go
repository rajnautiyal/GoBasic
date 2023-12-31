package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var wtGp sync.WaitGroup

func TestConcurrency() {
	fmt.Println("printing foo and baar")
	wtGp.Add(2)
	go printFooNumber()
	go printBaarNumber()
	wtGp.Wait()

}

func printFooNumber() {
	for i := 0; i < 10; i++ {
		fmt.Println("this is foo : ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wtGp.Done()
}

func printBaarNumber() {
	for i := 0; i < 10; i++ {
		fmt.Println("this is baar : ", i)
		time.Sleep(time.Duration(7 * time.Millisecond))
	}

	wtGp.Done()
}
