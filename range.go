package main

import (
	"fmt"
)

func main() {
	numeros := []int{2,4,6}
	suma := 0
	for _, numero := range numeros{
		suma += numero
		fmt.Printf("Suma:", suma)
		for i, numero := range numeros{
			if numero == 3
				fmt.Println("index:", i)
		}
	}
}

