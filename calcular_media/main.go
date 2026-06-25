package main

import "fmt"

var n1, n2, n3, n4, media float64

func main() {
	fmt.Println("Digite a primeira nota: ")
	fmt.Scanln(&n1)

	fmt.Println("Digite a segunda nota: ")
	fmt.Scanln(&n2)

	fmt.Println("Digite a terceira nota: ")
	fmt.Scanln(&n3)

	fmt.Println("Digite a quarta nota: ")
	fmt.Scanln(&n4)

	media = (n1 + n2 + n3 + n4) / 4

	if media >= 7 {
		fmt.Printf("A média é %.2f. Aluno aprovado!", media)
	}else {
		fmt.Printf("A média é %.2f. Aluno reprovado!", media)
	}
}
