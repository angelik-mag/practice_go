package main

import (
	"fmt"
	"runtime"

	"github.com/angelik-mag/practice_go/ejercicios"
	"github.com/angelik-mag/practice_go/variables"
)

// Función principal de GO
func main() {
	estado, texto := variables.ConvertirTexto(752)

	fmt.Println("estado = ", estado)
	fmt.Println("texto = ", texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("No es windows")
	} else {
		fmt.Println("Es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Es Linux")
	case "darwin":
		fmt.Println("Es darwin")
	default:
		fmt.Printf("%s \n", os)

	}

	valInt, respuesta := ejercicios.VerificarValor("30")

	fmt.Println("Numero = ", valInt)
	fmt.Println("Verificacion =", respuesta)
}
