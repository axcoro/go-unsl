package main
import "fmt"

func main() {
	a, b := 0, "b"
	fmt.Println(a, b)
	c, d := palabras()
	fmt.Println(c, d)
	e, _ := palabras()
	fmt.Println(e)
}

func palabras() (string, error) {
	return "valor de retorno", fmt.Errorf("Error")
}
