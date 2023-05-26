package ejercicios

import (
	"strconv"
)

func ConversorNumerico(texto string) (int, string) {
	entero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Se produjo un error: " + err.Error()
	}
	if entero > 100 {
		return entero, "MAYOR A 100"
	} else {
		return entero, "MENOR A 100"
	}
}
