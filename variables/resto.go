package variables

import (
	"fmt"
	"strconv"
	"time"
)

var info_local string //variable con scope en este archivo nada más
var Nombre string     //al tener el nombre en mayúscula el alcance es desde cualquier archivo del paquete
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.65
	Fecha = time.Now()

	fmt.Println("Resto de variables:")
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
	fmt.Println()
}

func ConviertoAtexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero)
	return true, texto
}
