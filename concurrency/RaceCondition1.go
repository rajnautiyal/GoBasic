package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// counter is a variable incremented by all goroutines.
var counter int
var syWt1 sync.WaitGroup

var mutex sync.Mutex

func TestRaceCondtion() {
	syWt1.Add(2)
	go incremtor("foo")
	go incremtor("baar")

	syWt1.Wait()
	fmt.Println("Final Counter :", counter)
}

func incremtor(s string) {
	for i := 0; i < 20; i++ {
		mutex.Lock()
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter :", counter)
		mutex.Unlock()
	}
	syWt1.Done()
}
