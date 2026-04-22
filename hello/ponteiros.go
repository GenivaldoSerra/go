package main

import "fmt"


func main() {
	var inteiro = 45
	var ponteiro *int = &inteiro

	fmt.Println("Valor do inteiro: ", inteiro)
	fmt.Println("Endereço do inteiro: ", &inteiro)
	fmt.Println("Valor do ponteiro: ", ponteiro)
	fmt.Println("Valor apontado pelo ponteiro: ", *ponteiro)
}