package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	time.Sleep(5 * time.Second)

	end := time.Now()

	duration := end.Sub(start)
	fmt.Println("Duration: ", duration.Milliseconds())
}
