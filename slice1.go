package main

import (
	"fmt"
)

func main() {
	// Crear un slice (lista) de números decimales tipo float64
	notas := []float64{4.5, 3.8, 4.0, 5.0, 3.6}

	// Imprimir el slice completo
	fmt.Println("Notas:", notas)

	// Recorrer el slice con un bucle for
	fmt.Println("Lista de notas una por una:")
	for i, valor := range notas {
		fmt.Printf("Posición %d → %.2f\n", i, valor)
	}

	// Calcular el promedio
	var suma float64
	for _, valor := range notas {
		suma += valor
	}
	promedio := suma / float64(len(notas))

	fmt.Printf("\nPromedio final: %.2f\n", promedio)
}
