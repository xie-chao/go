// regular-expressions.go
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var pn = fmt.Println

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	pn(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	pn(r.MatchString("prach"))
	pn(r.FindString("peach punch"))
	pn(r.FindStringIndex("peach punch"))
	pn(r.FindStringSubmatch("peach punch"))
	pn(r.FindStringSubmatchIndex("peach punch"))
	pn(r.FindAllString("peach punch pinch", -1))
	pn(r.FindAllStringSubmatch("peach punch pinch", -1))
	pn(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	pn(r.FindAllString("peach punch pinch", 2))
	pn(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	pn(r)
	pn(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	pn(string(out))
}
