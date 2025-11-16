package main

import "fmt"

func swap(a, b string) (string, string) {
	return a, b
}

func main() {
	a, b := swap("Hola", "Elkin")
	fmt.Println(a, b)
}
