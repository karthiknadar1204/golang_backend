package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", i)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	//start 3 worker go routines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers finished")
}
