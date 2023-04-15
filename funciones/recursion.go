package funciones

import "fmt"

func CalcularExponente(valor int) {

	if valor > 1000 {
		return
	}
	fmt.Println(valor)
	CalcularExponente(valor * 2)
}
