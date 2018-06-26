package main

const flag bool = true // constants

var (
	a int = 1                       // 'int' es redundante
	b     = []string{"a", "b", "c"} // declaracion inline de un array de strings
)

func main() {
	var b bool           // valor por defecto es false
	var m map[string]int // map[<tipo de la key>]<tipo del contenido>

	// make es "similar" al new de java
	m = make(map[string]int) // make(tipo[,tamaño])

	f := make([]string, 10) // declaracion inline de una nueva variable
}
