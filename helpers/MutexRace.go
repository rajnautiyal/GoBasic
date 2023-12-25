package helpers

// creating the race the condition where the counter is printing different values

//to check wether program is generating the raice coding you need to run "go run -race main.go"
import (
	"fmt"
	"runtime"
	"sync"
)

func MutexRaceCodititon() {
	counter := 0
	const gs = 50
	var wg sync.WaitGroup
	wg.Add(gs)

	fmt.Println("Number of CPU", runtime.NumCPU())
	fmt.Println("CPUS : ", runtime.NumGoroutine())
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutine : ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("NumGoroutine : ", runtime.NumGoroutine())
	fmt.Println("counter : ", counter)

}
