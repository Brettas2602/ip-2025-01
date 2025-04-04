package main

import "fmt"

func main() {
	var num int

	fmt.Println("Digite um número")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Printf("O número %v é positivo\n", num)
	}else if num < 0 {
		fmt.Printf("O número %v é negativo\n", num)
	}else {
		fmt.Printf("O número %v é nulo\n", num)
	}
}