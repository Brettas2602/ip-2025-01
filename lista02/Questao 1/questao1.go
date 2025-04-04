package main

import "fmt"

func main()  {
	var num int

	fmt.Println("Digite um número")
	fmt.Scan(&num)

	if num % 2 == 0 {
		fmt.Printf("O número %v é par\n", num)
	}else {
		fmt.Printf("O número %v é ímpar\n", num)
	}
}