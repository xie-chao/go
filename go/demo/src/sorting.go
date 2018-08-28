package main

import (
	"fmt"
	//	"sort"
)

type obj struct {
	test int
}

func change(arr []string) {
	arr[0] = "x"
	fmt.Println(arr)
}

func change1(o obj) {
	o.test = 8
	fmt.Println(o)
}

func main() {
	strs := []string{"c", "a", "e", "b", "d"}
	//	sort.Strings(strs)
	change(strs)
	fmt.Println(strs)
	o := obj{test: 2}
	fmt.Println(o)
	change1(o)
	fmt.Println(o)
}
