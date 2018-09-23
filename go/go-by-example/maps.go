package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("map:", m)

	delete(m, "k2")
	fmt.Println("map:", m)

	v, prs := m["k2"]
	fmt.Println("v:", v)
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	j := map[string]string{"foo:1": "1", "bar": "2"}
	fmt.Println("map:", j)
}
