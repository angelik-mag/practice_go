package main

import (
	d "github.com/angelik-mag/practice_go/defer_panic"
)

// Función principal de GO
func main() {

	/*
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

				valInt, respuesta := ejercicios.VerificarValor("334dd")

				fmt.Println("Numero = ", valInt)
				fmt.Println("Verificacion =", respuesta)
				teclado.IngresoNumeros()
				iteraciones.Iter()

			//fmt.Println(ejercicios.GenerarTabla())
			files.SumarTabla()
			arreglos_slices.Capacidad()
			mapas.MostrarMapas()

		Pedro := new(modelos.Hombre)
		Ana := new(modelos.Mujer)
		e.HumanosRespirando(Pedro)
		e.HumanosRespirando(Ana)

	*/

	d.VerDefer()
}
