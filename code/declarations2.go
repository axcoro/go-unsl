package main

import (
	"fmt"
)

type oracion string

func main() {

	c := oracion("hola mundo!") // tipo(valor)
	fmt.Printf("oración: %v, string: %v", c, string(c)) // casteo al tipo base

}
