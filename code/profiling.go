package main

import (
	"log"
	"os"
	"regexp"
	"runtime/pprof"
	"strings"

	"github.com/bxcodec/faker"
)

//START OMIT
type DatosFalsos struct {
	Email string `faker:"email"`
}

var Datos = []DatosFalsos{}

func main() {
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()
	validar()
}

func validar() {
	for _, value := range Datos {
		validEmail := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,5}$`)
		value.Email = strings.ToLower(value.Email)
		if !validEmail.MatchString(value.Email) {
			log.Panicf("No es valido %v", value.Email)
		}
	}
}

//END OMIT
func init() {
	a := DatosFalsos{}
	for i := 0; i < 500000; i++ {
		faker.FakeData(&a)
		Datos = append(Datos, a)
	}
}
