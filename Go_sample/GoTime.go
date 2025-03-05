package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	for i := 0; i < 1000000; i++ {
		fmt.Println(i)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Time took: %v\n", duration)
}
