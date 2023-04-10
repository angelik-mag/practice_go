package teclado

import (
	"bufio" // libreria del io -lectura por teclado, archivos
	"fmt"
	"os" //sistema operativo
	"strconv"
)

var numero_uno int
var numero_dos int
var leyenda string
var err error

func IngresoNumeros() {

	scanner := bufio.NewScanner(os.Stdin) // Scanner de teclado, tambien permite trabajar con archivos y otros

	fmt.Println("Ingrese número 1 : ")

	if scanner.Scan() {
		numero_uno, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato no sirve : " + err.Error())
		}
	}

	fmt.Println("Ingrese número 2 : ")
	if scanner.Scan() {
		numero_dos, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato no sirve : " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda : ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero_uno*numero_dos)
}
