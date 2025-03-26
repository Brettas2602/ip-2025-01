package main

import ("fmt")

func main()  {
	var numero int
	var soma float64

	fmt.Scan(&numero)

	for i := 1; i <= numero; i++ {
		soma += (1 / float64(i))
	}

	fmt.Printf("%.6f\n", soma)
}