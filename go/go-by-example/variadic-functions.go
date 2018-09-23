package main

import (
	"fmt"
)

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(2, 2, 3))

	nums := [4]int{2, 3, 4, 5}
	fmt.Println(sum(nums[:]...))
}
