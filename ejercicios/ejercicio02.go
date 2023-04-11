package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func GenerarTabla() string{
	
	var texto string
	scanner := bufio.NewScanner(os.Stdin) // Scanner de teclado, tambien permite trabajar con archivos y otros

	for {
		fmt.Println("Digite un valor")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El valor no es válido, digite nuevamente")
			} else {
				break
			}
		}
	}

	for i := 0; i <= 10; i++ {
		texto += fmt.Sprintf(" %d  x %d = %d \n ", numero, i, numero*i)
	}

	return texto
}
