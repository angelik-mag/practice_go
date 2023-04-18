package defer_panic

import (
	"fmt"
)

func VerDefer() {

	fmt.Println("Primero mensaje ")
	defer fmt.Println("mensaje final mensaje ")

	fmt.Println("Segundo mensaje ")

}
