package main

import (
	"fmt"

	"github.com/anibalvenezia/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoAtexto(384)
	fmt.Println(estado)
	fmt.Println(texto)
}
