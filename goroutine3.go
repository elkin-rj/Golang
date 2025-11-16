package main

import (
	"fmt"
)

func main() {
	suma := 0
	resta := 0

	for i := 1; i <= 10; i++ {

		suma = suma + i
		resta = resta - i

		fmt.Println("La suma es:", suma)
		fmt.Println("La resta es", resta)

	}

}
