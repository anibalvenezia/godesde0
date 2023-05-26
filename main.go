package main

import (
	"fmt"

	"github.com/anibalvenezia/godesde0/ejercicios"
)

//"fmt"
//"runtime"

func main() {
	/*
		variables.MuestroEnteros()
		variables.RestoVariables()
		estado, texto := variables.ConviertoAtexto(384)
		fmt.Println(estado)
		fmt.Println(texto)
		// Aprendemos sintaxis de IF en GO
		os := runtime.GOOS
		if os == "linux" {

		} else {
			fmt.Println("Su OS no es Linux, es ", os)
		}
		// También puedo hacer la asignación dentro del IF
		if goarch := runtime.GOARCH; goarch == "amd64" {
			fmt.Println("Arquitectura ", goarch)
		}
		// Aprendemos sentencia SWITCH
		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es Linux")
		case "darwin":
			fmt.Println("Esto es Darwin")
		default:
			fmt.Printf("%s \n", os)
		}
	*/

	//EJERCICIO 01
	numero, texto := ejercicios.ConversorNumerico("227")
	fmt.Println(numero, texto)

}
