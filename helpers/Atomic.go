package helpers

// creating the race the condition where the counter is printing different values

//to check wether program is generating the raice coding you need to run "go run -race main.go"
import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func Atomic() {
	var counter int64
	const gs = 50
	var wg sync.WaitGroup
	wg.Add(gs)

	fmt.Println("Number of CPU", runtime.NumCPU())
	fmt.Println("CPUS : ", runtime.NumGoroutine())
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)

			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Goroutine : ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("NumGoroutine : ", runtime.NumGoroutine())
	fmt.Println("counter : ", counter)

}
