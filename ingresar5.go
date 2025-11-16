package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingresa tu fecha de nacimiento (dd/mm/aaaa): ")
	entrada, _ := reader.ReadString('\n')
	entrada = strings.TrimSpace(entrada)

	// Intentar analizar la fecha con el formato exacto
	fechaNacimiento, err := time.Parse("02/01/2006", entrada)
	if err != nil {
		fmt.Println("⚠️ Error: Formato inválido. Debes usar el formato dd/mm/aaaa (por ejemplo 04/11/1976).")
		return
	}

	// Obtener la fecha actual
	hoy := time.Now()

	// Verificar que la fecha no sea futura
	if fechaNacimiento.After(hoy) {
		fmt.Println("⚠️ Error: La fecha de nacimiento no puede ser en el futuro.")
		return
	}

	// Calcular la edad
	edad := hoy.Year() - fechaNacimiento.Year()

	// Ajustar si aún no ha cumplido años este año
	if hoy.Month() < fechaNacimiento.Month() ||
		(hoy.Month() == fechaNacimiento.Month() && hoy.Day() < fechaNacimiento.Day()) {
		edad--
	}

	fmt.Printf("✅ Fecha válida: %s\n", fechaNacimiento.Format("02 de January de 2006"))
	fmt.Printf("Tienes %d años.\n", edad)
}
