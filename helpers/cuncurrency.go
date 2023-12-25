package helpers

import (
	"fmt"
	"runtime"
	"sync"
)

func Concurrency() {
	for i := 0; i < 10; i++ {
		fmt.Println("this is the ", i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println("this is the ", i)
	}
	//fmt.Println(runtime.CPUProfile())

}

func OSDetails(Wg *sync.WaitGroup) {
	fmt.Println("inisde here")
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("GOARCH: ", runtime.GOARCH)
	fmt.Println("Goroutine : ", runtime.NumGoroutine())
	Wg.Done()
}
