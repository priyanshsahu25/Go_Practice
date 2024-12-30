package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}
type Square struct {
	side float64
}
type rectangle struct {
	length float64
	width  float64
}

func (rec rectangle) Area() float64 {

	return rec.length * rec.width
}

func (cir Circle) Area() float64 {
	return 3.14 * cir.radius * cir.radius
}

func (sq Square) Area() float64 {
	return sq.side * sq.side
}
func main() {

	rec := rectangle{2, 3}
	cir := Circle{2}
	sq := Square{5}
	fmt.Println(rec.Area())
	fmt.Println(cir.Area())
	fmt.Println(sq.Area())

}
