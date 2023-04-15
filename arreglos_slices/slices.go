package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 8, 0, 0, 0, 0, 89}

func MostrarSlice() {

	fmt.Println(arreglo)

	porcion := arreglo[3:] //slice creado desde un vector desde la posición 3

	porcion2 := arreglo[:5] //slice creado desde la pocisión 0 hasta la 5

	porcion3 := arreglo[3:7] // slice creado desde la pocisión 3 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad() {

	elementos := make([]int, 5, 20)

	nums := make([]int, 0, 0)

	for i := 0; i < 130; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\n nums: Largo %d, Capacidad %d ", len(nums), cap(nums))
	fmt.Printf("\n elementos: Largo %d, Capacidad %d ", len(elementos), cap(elementos))

}
