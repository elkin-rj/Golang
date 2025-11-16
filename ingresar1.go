package main

import "fmt"

func main() {
	var nombre, apellido string

	fmt.Print("Digite el nombre: ")
	fmt.Scanln(&nombre)

	fmt.Print("Digite el apellido: ")
	fmt.Scanln(&apellido)
	fmt.Println("Hola", nombre, apellido)

}
