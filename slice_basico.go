package main

import (
	"fmt"
)

func main() {

	notas := []float64{5, 9, 8, 6, 3}

	fmt.Println("Notas:", notas)
	fmt.Println(notas[3])
	valor := notas[3]

	fmt.Println(valor + 6)

}
