package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Tenemos %g inconvenientes técnicos. \n", math.Sqrt(7))  // muestra la raiz cuadrada de 7, en formato float64
	fmt.Printf("Tenemos %.1f inconvenientes técnicos.\n", math.Sqrt(7)) // muestra la raiz cuadrada de 7, con un solo decimal.
}
