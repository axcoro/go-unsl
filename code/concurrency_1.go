package main

import (
	"fmt"
)

func main() {

	go func() {
		fmt.Print("hago un print")
	}()

}
