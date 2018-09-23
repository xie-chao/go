// number-parsing
package main

import (
	"fmt"
	"strconv"
)

var pn = fmt.Println

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	pn(f)
	i, _ := strconv.ParseInt("1234", 0, 64)
	pn(i)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	pn(d)
	u, _ := strconv.ParseUint("789", 0, 64)
	pn(u)
	k, _ := strconv.Atoi("135")
	pn(k)
	_, e := strconv.Atoi("wtf")
	pn(e)
}
