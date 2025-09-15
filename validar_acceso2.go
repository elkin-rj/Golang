package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	fmt.Print("Ingrese la clave: ")

	// Lee la contraseña desde la entrada estándar sin mostrarla (sin eco)
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println() // imprime salto de línea después de que el usuario presione Enter
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error leyendo la clave:", err)
		return
	}

	password := string(bytePassword)

	if password == "123" {
		fmt.Println("Acceso Ok")
	} else {
		fmt.Println("Acceso denegado")
	}
}
