// url-parsing
package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	pn := fmt.Println
	s := "postgres://user:pass@host.com:5432/path?k=v&k=1#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	pn(u.Scheme)
	pn(u.User)
	pn(u.User.Username())
	pass, _ := u.User.Password()
	pn(pass)

	pn(u.Host)
	h := strings.Split(u.Host, ":")
	pn(h[0])
	pn(h[1])

	pn(u.Path)
	pn(u.Fragment)
	pn(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	pn(m)
	pn(m["k"])
}
