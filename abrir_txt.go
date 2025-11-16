package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("datos.txt")
	if err != nil {
		fmt.Println("⚠️ Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
