// reading-file
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/tmp/defer.txt")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("/tmp/defer.txt")
	check(err)

	b1 := make([]byte, 2)
	fmt.Println(b1)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
}
