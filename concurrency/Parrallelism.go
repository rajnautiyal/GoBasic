package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wtgroup2 sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func TestParralism() {
	wtgroup2.Add(2)
	printEventNumber()
	printOddNumber()
	wtgroup2.Wait()
}
func printEventNumber() {
	for i := 0; i < 45; i++ {
		fmt.Println("this is foo : ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wtgroup2.Done()
}

func printOddNumber() {
	for i := 0; i < 45; i++ {
		fmt.Println("this is baar : ", i)
		time.Sleep(time.Duration(7 * time.Millisecond))
	}

	wtgroup2.Done()
}
