package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/angelik-mag/practice_go/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

// Pisa los archivos
func GrabarTabla() {

	var texto string = ejercicios.GenerarTabla()
	archivo, err := os.Create("./files/txt/tabla.txt")

	if err != nil {
		fmt.Println("hay error")
		return
	}

	// Graba en archivo el texto
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumarTabla() {

	var texto string = ejercicios.GenerarTabla()

	if !Append(fileName, texto) {
		fmt.Println("hay error en la concatenación")
	}
}

func Append(file string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error en el append " + err.Error())
		return false
	}

	// Escritura en el archivo
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error en la escritura " + err.Error())
		return false
	}

	arch.Close()

	return true
}

func LeerArchivo() {

	archivo, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error en la lectura " + err.Error())

	}

	fmt.Println(string(archivo))

}

func LeerArchivoDos() {

	archivo, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error en la lectura 2 " + err.Error())

	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
