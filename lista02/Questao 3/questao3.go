package main

import "fmt"

func main()  {
	var num1, num2, soma int

	fmt.Println("Digite dois números")
	fmt.Scan(&num1, &num2)

	soma = num1 + num2

	if soma > 20 {
		fmt.Printf("A soma acrescida de 8 é %v\n", soma + 8)
	}else {
		fmt.Printf("A soma subtraida de 5 é %v\n", soma - 5)
	}
}