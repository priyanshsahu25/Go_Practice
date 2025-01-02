package main

import (
	"fmt"
	"sync"
)

func counter(add *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	*add++
	mu.Unlock()
}

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex
	num := 1
	add := &num
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter(add, &wg, &mu)
	}
	wg.Wait()
	fmt.Println(num)
}
