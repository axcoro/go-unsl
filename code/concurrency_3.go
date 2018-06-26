package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("closure", i)
		}()
	}

	time.Sleep(1000 * time.Millisecond)
}
