package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	result := r.width * r.height
	return result
}

func (r rect) perim() int {
	result := (r.width + r.height) * 2

	r.width = 1
	return result
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("area:", r.perim())
	fmt.Println(r.width)

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println(r.width)

	fmt.Println("area:", rp.perim())
}
