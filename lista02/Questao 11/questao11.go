package main

import "fmt"

func main()  {
	var x, y float64

	fmt.Println("Digite um número")
	fmt.Scan(&x)

	y = 8 / (2 - x)

	fmt.Printf("f(x) = %.2f", y)
}