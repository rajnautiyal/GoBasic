package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// counter is a variable incremented by all goroutines.
var counterAtom int64
var syncWt sync.WaitGroup

func TestAtomicCondtion() {
	syncWt.Add(2)
	go incremtor3("foo")
	go incremtor3("baar")

	syncWt.Wait()
	fmt.Println("Final Counter :", counterAtom)
}

func incremtor3(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&counterAtom, 1)
		fmt.Println(s, i, "Counter :", counterAtom)

	}
	syncWt.Done()
}
