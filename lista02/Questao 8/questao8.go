package main

import "fmt"

func main()  {
	var num int

	fmt.Println("Digite um número")
	fmt.Scan(&num)

	if num > 20 && num < 90 {
		fmt.Printf("O número %v está entre 20 e 90\n", num)
	}else {
		fmt.Printf("O número %v não está entre 20 e 90\n", num)
	}
}