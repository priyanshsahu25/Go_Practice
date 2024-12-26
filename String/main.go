package main

import "fmt"

func sum[T any](arr []T) {

	for _, val := range arr {
		fmt.Println(val)
	}

}

func main() {
	// arr := []int{1, 2, 3, 4}
	str := []string{"hello", " Priyansh "}
	fmt.Println(str)
}
