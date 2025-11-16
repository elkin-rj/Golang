package main

import (
	"fmt"
	"time"
)

func main() {
	// Crea un ticker que emite un "tick" cada segundo
	ticker := time.NewTicker(1 * time.Second)

	// Crea un temporizador que se activará después de 10 segundos
	stop := time.NewTimer(10 * time.Second)

	// Bucle infinito controlado por un select
	for {
		select {
		case <-ticker.C:
			fmt.Println("Hola Elkin")
		case <-stop.C:
			fmt.Println("⏹️ Tiempo finalizado, se detiene el programa.")
			ticker.Stop() // Detenemos el ticker para liberar recursos
			return        // Salimos de la función main
		}
	}
}
