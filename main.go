package main

import (
	"fmt"
)

func main() {

	var ch1 chan int
	ch2 := make(chan int)
	fmt.Println(ch1, ch2)
}
