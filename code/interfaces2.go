package main
import "fmt"
type Animal interface {
  hablar() string
}
type Pato struct {
  nombre string
}
type Perro struct {
  nombre string
}
func(animal Pato) hablar() string {//lucas.hablar()
  return "cuac cuac"
}
func(animal Perro) hablar() string {//toby.hablar()
  return "guau guau"
}
func hablar(animal Animal) string {//Consultar al metodo de manera enmascarada
  return animal.hablar()
}
func main() {
  toby, lucas := Perro {nombre: "Toby"},Pato {nombre: "Lucas"}
  fmt.Println("Toby me habla:", hablar(toby))//Uso de la función
  fmt.Println("Lucas me habla:", lucas.hablar())//Uso de registro para llamar al método
}
