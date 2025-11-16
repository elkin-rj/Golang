package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var edad int
	reader := bufio.NewReader(os.Stdin)
	reader2 := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Ingresa tu nombre: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Validar que todos los caracteres sean letras o espacios
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
			fmt.Println("✅ Nombre válido:", input)
			break // Sale del bucle si la entrada es válida
		}
	}

	for {
		fmt.Print("Ingresa tu apellido: ")
		input2, _ := reader2.ReadString('\n')
		input2 = strings.TrimSpace(input2)

		// Validar que todos los caracteres sean letras o espacios
		esTexto2 := true
		for _, ch2 := range input2 {
			if !unicode.IsLetter(ch2) && ch2 != ' ' {
				esTexto2 = false
				break
			}
		}

		if !esTexto2 || input2 == "" {
			fmt.Println("⚠️ Error: Solo se permiten letras (sin números ni símbolos). Intenta de nuevo.\n")
		} else {
			fmt.Println("✅ Nombre válido:", input2)
			break // Sale del bucle si la entrada es válida
		}
	}

	fmt.Print("Digite su edad:")
	fmt.Scanln(&edad)
	fmt.Println("su edad es:", edad)

}
