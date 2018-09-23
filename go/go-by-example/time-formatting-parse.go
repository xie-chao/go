// time-formatting-parse.go
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(time.RFC3339, "2018-08-05T17:26:56+08:00")
	if e != nil {
		panic(e)
	}
	p(t1)
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02 15:04:05"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	if e != nil {
		panic(e)
	}
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
