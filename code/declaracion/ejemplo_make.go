package main

import "fmt"

func main() {
	intList := make([]int, 3)
	fmt.Println("[]int: ", intList)

	pointerList := make([]*string, 3)
	fmt.Println("[]*string: ",pointerList)

	stringChanel := make(chan []string, 0)
	fmt.Println("chan []string: ", stringChanel)

	stringMap := make(map[string]string, 0)
	fmt.Println("map[string]string: ", stringMap)
}
