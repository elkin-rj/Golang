package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Registrar un manejador (handler) para la ruta "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "¡Hola, mundo! Servidor Go en marcha.")
	})

	// 2. Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// Si hay un error al iniciar, se imprime y el programa termina.
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
