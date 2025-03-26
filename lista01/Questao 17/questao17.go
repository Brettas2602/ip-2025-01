package main

import ("fmt")

func main()  {
	var primeiroNum, segundoNum int

	fmt.Scan(&primeiroNum, &segundoNum)

	if primeiroNum % 2 != 0 {
		fmt.Print("O PRIMEIRO NUMERO NAO E PAR")
	} else {
		for i := 0; i < segundoNum; i++ {
			fmt.Print(primeiroNum, " ")
			primeiroNum += 2
		}
	}

	fmt.Println("")
}