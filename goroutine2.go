package main

import (
	"fmt"
	"time"
)

func saludar(nombre string) {
	for i := 1; i <= 3; i++ {
		fmt.Println("Hola", nombre, "-", i)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	go saludar("Elkin") // ✅ se ejecuta concurrentemente
	go saludar("Ana")   // ✅ también concurrente

	fmt.Println("Esperando que terminen las goroutines...")
	time.Sleep(3 * time.Second)
	fmt.Println("Programa terminado.")
}
