package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(1 * time.Second)
	for {
		fmt.Println("Hola Elkin")
		<-ticker.C
	}
}
