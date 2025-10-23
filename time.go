package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("La hora del sistema es:", t.Format("01:01 AM")) // Hora del sistema
	fmt.Println("La hora es:", time.Now())
	fmt.Println("La hora es:", time.Now().Year())
	fmt.Println("El minuto es:", t.Minute())
	fmt.Println("La hora es:", t.Hour())

	hora := time.Date(2025, time.October, 22, 14, 30, 0, 0, time.Local) // Hora fija o simulada
	fmt.Println("Hora creada manualmente:", hora)
}
