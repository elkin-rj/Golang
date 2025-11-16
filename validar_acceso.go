package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese la clave: ")

	// Lee hasta presionar Enter
	clave, _ := reader.ReadString('\n')
	// Elimina salto de línea y espacios
	clave = strings.TrimSpace(clave)

	if clave == "123" {
		fmt.Println("Acceso Ok")
	} else {
		fmt.Println("Acceso denegado")
	}
}
