package mapas

import "fmt"

func MostrarMapas() {

	paises := make(map[string]string)

	paises["Mexico"] = "D.F."
	paises["Colombia"] = "Bogota"

	// fmt.Println(paises)
	// fmt.Println(paises["Colombia"])

	campeonato := map[string]int{
		"Barcelona": 39,

		"Real Madrid": 48,
		"Boca Junior": 41,
		"America":     51,
		"Liver":       78,
	}

	fmt.Println(campeonato)

	// range es un especie de foreche
	for equipo, puntaje := range campeonato {

		fmt.Printf("\n El equipo %s tiene un puntale de %d ", equipo, puntaje)
	}

	// eliminar de un mapa
	delete(campeonato, "Liver")
	fmt.Println(campeonato)

	// range es un especie de foreche
	for equipo, puntaje := range campeonato {

		fmt.Printf("\n El equipo %s tiene un puntale de %d ", equipo, puntaje)
	}

	puntaje, existe := campeonato["Juventus"]

	fmt.Printf("\n El puntaje es %d y el equipo existe = %t ", puntaje, existe)
}
