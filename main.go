package main

import "fmt"

func main() {
	var i interface{} = "Hello"

	switch v := i.(type) {
	case int:
		fmt.Println("Intefer", v)
	case string:
		fmt.Println("STring", v)
	default:
		fmt.Println("NO DATA TYPE")
	}
}
