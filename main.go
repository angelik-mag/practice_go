package main

import (
	"fmt"

	"github.com/angelik-mag/practice_go/variables"
)

// Función principal de GO
func main() {
	estado, texto := variables.ConvertirTexto(752)

	fmt.Println("estado = ", estado)
	fmt.Println("texto = ", texto)
}
