package main

import "fmt"

func main() {
	lista := [4]int{2, 2, 3, 2}
	produto := 1

	for i := 0; i < len(lista); i++ {
		produto = produto * lista[i]
	}

	fmt.Print(produto)
}
