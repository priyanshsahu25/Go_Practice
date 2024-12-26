package main

import "fmt"

func show(arr []any) {
	fmt.Println(arr)
}

type Person struct {
	name   string
	gender any
	call   any
}

func sum[T any](arr []T) {
	for _, val := range arr {
		fmt.Println(val)
	}
}

func infr(arr []interface{}) {
	fmt.Println(arr)
}

func main() {
	arrstr := []string{"huhuh", "jiii"}
	arrint := []int{1, 2}
	sum(arrstr)
	sum(arrint)
	newarr := []interface{}{"hello", 1}
	fmt.Println(newarr)

	mpp := make(map[any]any)
	mpp["hello"] = 1
	mpp[22] = 11


	Priyansh := Person{"Sahu", 1, "jhjh"}
	Jivesh := Person{"jjj", "jbjbn", "njnjn"}
	fmt.Println(Priyansh, Jivesh)

	mpp1 := make(map[any]any)
	mpp1["ID"] = 1
	mpp1[123] = "Hello"
	fmt.Println(mpp1)
}
