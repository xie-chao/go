package main

import (
	"fmt"
)

func plus(a int, b int) (int, int, int) {
	return a, b, a + b
}

func main() {
	n1, n2, res := plus(1, 2)
	fmt.Printf("%d + %d = %d\n", n1, n2, res)
}
