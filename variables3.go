package main

import (
	"fmt"
)

func main() {

	var A byte = 'A'
	var a byte = 'a'
	fmt.Println(A)
	fmt.Println(a)

	var suma int = 10
	fmt.Println(suma > 10 || suma == 10)

	var numero int
	switch {
	case numero > 0:
		fmt.Println("Es positivo")
	case numero < 0:
		fmt.Println("Es negativo")
	case numero == 0:
		fmt.Println("Es cero")
	}
}
