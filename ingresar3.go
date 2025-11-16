package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// --- Validar nombre ---
	var nombre string
	for {
		fmt.Print("Ingresa tu nombre: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		esTexto := true
		for _, ch := range input {
			if !unicode.IsLetter(ch) && ch != ' ' {
				esTexto = false
				break
			}
		}

		if !esTexto || input == "" {
			fmt.Println("⚠️ Error: Solo se permiten letras (sin números ni símbolos). Intenta de nuevo.\n")
		} else {
			nombre = input
			break
		}
	}

	// --- Validar apellido ---
	var apellido string
	for {
		fmt.Print("Ingresa tu apellido: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		esTexto := true
		for _, ch := range input {
			if !unicode.IsLetter(ch) && ch != ' ' {
				esTexto = false
				break
			}
		}

		if !esTexto || input == "" {
			fmt.Println("⚠️ Error: Solo se permiten letras (sin números ni símbolos). Intenta de nuevo.\n")
		} else {
			apellido = input
			break
		}
	}

	// --- Validar edad ---
	var edad int
	for {
		fmt.Print("Ingresa tu edad: ")
		_, err := fmt.Scanln(&edad) // lee e intenta convertir a número entero

		if err != nil {
			fmt.Println("⚠️ Error: Debes ingresar un número entero válido.")
			// limpiar el buffer del teclado antes de volver a pedir el dato
			var limpiar string
			fmt.Scanln(&limpiar)
		} else if edad <= 0 {
			fmt.Println("⚠️ Error: La edad debe ser mayor que cero.")
		} else {
			break
		}
	}

	fmt.Println("\n✅ Datos ingresados correctamente:")
	fmt.Println("Nombre:", nombre)
	fmt.Println("Apellido:", apellido)
	fmt.Println("Edad:", edad)
}
