// epoch.go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	fmt.Println("secs:", secs)
	millis := nanos / 1000000
	fmt.Println("millis:", millis)
	fmt.Println("nanos:", nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
