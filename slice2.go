package main

import (
	"fmt"
)

func main() {

	var slice1 [6]int
	fmt.Println(slice1) // Resultado [0 0 0 0 0 0]

	// Crear un slice (lista) de números decimales tipo float64
	notas := []float64{4, 3.5, 5.8, 6.7}
	fmt.Println("Las notas son:", notas)
	fmt.Println("este es el valor en la posición 2:", notas[2])
	notas[3] = 9
	notas = append(notas, 16)
	fmt.Println("este es el valor en la posición 3:", notas[3])
	fmt.Println("Tamaño del slice:", len(notas))
	fmt.Println("Las notas nuevas son:", notas)
	notas2 := notas[0:2]
	fmt.Println("Nuevo slice:", notas2)

	placa_vehi := []string{
		"abc23", 
		"jkh953"
	}
	fmt.Println("Placas vehículos:", placa_vehi)
}
