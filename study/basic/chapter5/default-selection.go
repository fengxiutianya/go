package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Microsecond)
	boom := time.Tick(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("      .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}