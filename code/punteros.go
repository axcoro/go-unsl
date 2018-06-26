package main

import "fmt"

func main() {
	edad := 25
	var puntero *int //Declaracion de un puntero
	puntero = &edad  //Asigno a puntero la posicion de memoria de edad

	fmt.Println("puntero: ", puntero)
	fmt.Println("&edad: ", &edad)
	fmt.Println("&puntero: ", &puntero)
	fmt.Println("*puntero: ", *puntero)
}
