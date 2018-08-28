// json.go
package main

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

var pn = fmt.Println

func main() {
	bolB, _ := json.Marshal(true)
	pn(string(bolB))

	intB, _ := json.Marshal(1)
	pn(string(intB))

	fltB, _ := json.Marshal(3.14)
	pn(string(fltB))

	strB, _ := json.Marshal("gopher")
	pn(string(strB))

	slcD := []string{"apple", "prach", "pear"}
	slcB, _ := json.Marshal(slcD)
	pn(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	pn(string(mapB))

	res1D := &Response1{
		Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	pn(string(res1B))

	res2D := &Response2{
		Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	pn(string(res2B))

	byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	pn(dat)

}
