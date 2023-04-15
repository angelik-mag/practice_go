package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 2, 45, 6, 8}
var matriz [3][4]int

func MostrarArreglo() {
	// formas de asignar valor

	tabla[7] = 45
	tabla[5] = 36
	//fmt.Println(tabla)

	fmt.Println(matriz)

	matriz[1][3] = 25
	fmt.Println(matriz)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
}
