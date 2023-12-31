package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// counter is a variable incremented by all goroutines.
var counter1 int
var syWt2 sync.WaitGroup

var mutex1 sync.Mutex

func TestFixRaceCondtion1() {
	syWt2.Add(2)
	go incremtor1("foo")
	go incremtor1("baar")

	syWt2.Wait()
	fmt.Println("Final Counter :", counter1)
}

func incremtor1(s string) {
	for i := 0; i < 20; i++ {
		mutex1.Lock()
		x := counter1
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter1 = x
		fmt.Println(s, i, "Counter :", counter1)
		mutex1.Unlock()
	}
	syWt2.Done()
}
