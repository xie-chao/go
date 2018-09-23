package main

import (
	"fmt"
	sf "strings"
)

var p = fmt.Println

func main() {
	p("Contains:", sf.Contains("test", "es"))
	p("Count:", sf.Count("test", "t"))
	p("HasPrefix:", sf.HasPrefix("test", "te"))
	p("HasSuffix:", sf.HasSuffix("test", "st"))
	p("Index:", sf.Index("test", "s"))
	p("Join:", sf.Join([]string{"a", "b"}, "-"))
	p("Repeat:", sf.Repeat("a", 5))
	p("Replace:", sf.Replace("foo", "o", "0", -1))
	p("Replace:", sf.Replace("foo", "o", "0", 1))
	p("Replace:", sf.Replace("fffooo", "o", "0", 2))
	p("Split:", sf.Split("a-b-c-d-e", "-"))
	p("ToLower:", sf.ToLower("TEST"))
	p("ToUpper:", sf.ToUpper("test"))

	p("Len:", len("hello"))
	p("Char:", "hello"[1])
}
