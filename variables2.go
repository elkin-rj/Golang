package main

import (
	"fmt"
)

func main() {
	var suma float64
	suma = 5.6

	fmt.Printf("La estructura del número es: %T, valor: %.1f\n", suma, suma)
	mostrarNombre()
}

func mostrarNombre() {
	fmt.Println("\n")
	fmt.Println(" ███████ ██      ██   ██ ██ ███  ")
	fmt.Println(" ██      ██      ██  ██  ██ ████ ")
	fmt.Println(" █████   ██      █████   ██ ██ ██")
	fmt.Println(" ██      ██      ██  ██  ██ ██  █")
	fmt.Println(" ███████ ███████ ██   ██ ██ ██  █")
	fmt.Println("\n        Developer: Elkin")
}
