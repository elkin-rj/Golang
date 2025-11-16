package main

import (
	"fmt"
	"time"
)

func main() {
	// Pedir la fecha de nacimiento
	var dia, mes, anio int
	fmt.Print("Ingresa tu fecha de nacimiento (día mes año): ")
	fmt.Scan(&dia, &mes, &anio)

	// Crear objeto de tipo time.Time
	fechaNacimiento := time.Date(anio, time.Month(mes), dia, 0, 0, 0, 0, time.Local)

	// Obtener la fecha actual
	hoy := time.Now()

	// Calcular la edad
	edad := hoy.Year() - fechaNacimiento.Year()

	// Ajustar si aún no ha cumplido años este año
	if hoy.Month() < fechaNacimiento.Month() ||
		(hoy.Month() == fechaNacimiento.Month() && hoy.Day() < fechaNacimiento.Day()) {
		edad--
	}

	fmt.Printf("Tienes %d años.\n", edad)
}
