package main

import "fmt"

func main()  {
	var num int

	fmt.Println("Digite um número")
	fmt.Scan(&num)

	if num % 5 == 0 {
		fmt.Printf("O número %v é divisível por 5\n", num)
	}else {
		fmt.Printf("O número %v não é divisível por 5\n", num)
	}
}