package ejercicios

import (
	"fmt"
	"strconv"
)

func VerificarValor(numero string) (int, string) {

	valInt, err := strconv.Atoi(numero)
	fmt.Println("Error = ", err)
	if valInt > 100 {
		return valInt, "Es mayor de 100"
	} else {
		return valInt, "Es menor de 100"
	}
}
