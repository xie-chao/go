package main

import (
	"fmt"
	"os"
	//	"os"
)

type point struct {
	x, y int
}

var pf = fmt.Printf

func main() {
	p := point{1, 2}
	pf("%v\n", p)
	pf("%+v\n", p)
	pf("%#v\n", p)
	pf("%T\n", p)
	pf("%t\n", true)
	pf("%d\n", 123)
	pf("%b\n", 14)
	pf("%c\n", 33)
	pf("%x\n", 456)
	pf("%f\n", 78.9)
	pf("%e\n", 123400000.0)
	pf("%E\n", 123400000.0)
	pf("%x\n", "hex this")
	pf("%p\n", &p)
	pf("|%6d|%6d|\n", 12, 345)
	pf("|%6.2f|%6.2f|\n", 1.2, 333.45)
	pf("|%6.2f|%6.2f|\n", 1.2, 3.45345)
	pf("|%-6.2f|%-6.2f|\n", 1.2, 3.45345)
	pf("|%6s|%6s|\n", "foo", "bar")
	pf("|%-6s|%-6s|\n", "foo", "bar")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
	fmt.Println(os.TempDir())
}
