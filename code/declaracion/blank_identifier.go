package main

import "fmt"
import _ "strings" // un import no utilizado es un error de compilacion

func main() {
	numeros := []string{"cero", "uno", "dos", "tres"}
	for _, valor := range numeros { // ignoramos el indice del valor
		fmt.Println(valor)
	}
}
