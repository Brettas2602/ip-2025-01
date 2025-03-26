package main

import ("fmt")

type Temperatura struct {
	fahrenheit float64
	celsius float64
}

func main()  {
	var numeroDeLeituras int
	var temperaturas = []Temperatura{}

	fmt.Scan(&numeroDeLeituras)

	for i := 0; i < numeroDeLeituras; i++ {
		var temp Temperatura
		fmt.Scan(&temp.fahrenheit)
		temp.celsius = (5 * (temp.fahrenheit - 32)) / 9

		temperaturas = append(temperaturas, temp)
	}

	for _, value := range temperaturas {
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", value.fahrenheit, value.celsius)
	}
}