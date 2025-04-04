package main

import "fmt"

func main()  {
	var A, B int

	fmt.Println("Digite um número A e um número B")
	fmt.Scan(&A, &B)

	if A % B == 0 {
		fmt.Printf("O número %v é divisível por %v\n", A, B)
	}else {
		fmt.Printf("O número %v não é divisível por %v\n", A, B)
	}
}