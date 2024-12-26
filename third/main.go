package main

import (
	"fmt"
	"sync"
	"time"
)

func printNum(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go printNum(&wg, &mu)
		time.Sleep(3*time.Second)
	}
	wg.Wait()
	fmt.Println("DONE")

}
