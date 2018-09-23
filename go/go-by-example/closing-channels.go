package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received jbos", j)
			} else {
				fmt.Println("received all jbos")
				done <- true
				fmt.Println("done...")
				return
			}
		}
	}()

	for j := 1; j < 10; j++ {
		jobs <- j
		fmt.Println("sent jobs", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
