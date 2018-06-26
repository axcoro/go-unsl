package main

import "fmt"

func main() {
	intList := make([]int, 3)
	fmt.Println(intList)

	pointerList := make([]*string, 3)
	fmt.Println(pointerList)

	stringChanle := make(chan []string, 0)
	fmt.Println(stringChanle)

	stringMap := make(map[string]string, 0)
	fmt.Println(stringMap)
}
