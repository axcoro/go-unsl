package main

import (
	"fmt"
)

func f(c chan<- string) {
	c <- "function"
}

func main() {
	ch := make(chan string)

	go func() { ch <- "closure" }()
	go f(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
