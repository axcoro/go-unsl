package controllers

import "github.com/UNSL/app/utils"

func main() {
	utils.UnNombre = 5  // OK
	utils.otroNombre = 3 // compile error
	utils.algunaFuncion() // compile error
}
