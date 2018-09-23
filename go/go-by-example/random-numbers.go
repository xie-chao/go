// random-numbers.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var p = fmt.Print
var pn = fmt.Println

func main() {
	p(rand.Intn(100), ",")
	pn(rand.Intn(100))

	pn(rand.Float64())
	p((rand.Float64()*5)+5, ",")
	pn((rand.Float64()*5)+5, ",")

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	p(r1.Intn(100), ",")
	pn(r1.Intn(100), ",")

	s2 := rand.NewSource(11)
	r2 := rand.New(s2)
	p(r2.Intn(100), ",")
	pn(r2.Intn(100), ",")

	printRand()
	printRand()
	printRand()

}

func printRand() {
	pn(rand.Intn(100))
}
