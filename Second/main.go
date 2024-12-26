package main

import (
	"fmt"
	"sync"
)

func sumEven(arr []int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	sum := 0
	for _, val := range arr {
		if val%2 == 0 {
			sum += val
		}
	}
	ch <- sum
}
func sumOdd(arr []int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	sum := 0
	for _, val := range arr {
		if val%2 != 0 {
			sum += val
		}
	}
	ch <- sum
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 2)
	arr := []int{1, 2, 4, 5, 6, 7, 8, 967, 89, 34, 56, 45}
	wg.Add(1)
	go sumEven(arr, &wg, ch)
	wg.Add(1)
	go sumOdd(arr, &wg, ch)
	wg.Wait()
	close(ch)
	even := <-ch
	odd := <-ch
	fmt.Println(even + odd)
}
