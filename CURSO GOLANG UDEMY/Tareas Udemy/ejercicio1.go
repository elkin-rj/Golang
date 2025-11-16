package main

import (
	"fmt"
)

func main() {

	var licencia bool
	var edad int
	fmt.Println("Digite tipo licencia: true o false:")
	fmt.Scanln(&licencia)
	fmt.Println("Digite la edad:")
	fmt.Scanln(&edad)
	if licencia == true && edad >= 15 {
		fmt.Println("Puede continuar")
	} else {
		fmt.Println("No puede continuar")
	}

}
