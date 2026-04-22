package main

import "fmt"

func pointerFuncao(x *int) {
	*a = 400
}

func main() {
	var x = 100
	fmt.Printf("Valor de x antes da função: %d\n", x)

	var pa *int = &x

	pointerFuncao(pa)
	fmt.Printf("Valor de x depois da função: %d\n", x)
}