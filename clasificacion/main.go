package main

import "fmt"

func main() {
	var n int
	fmt.Print("Ingresa un número entero positivo n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	switch {
	case n >= 1 && n <= 10:
		fmt.Println("Clasificación: Número pequeño")
	case n >= 11 && n <= 100:
		fmt.Println("Clasificación: Número mediano")
	case n > 100:
		fmt.Println("Clasificación: Número grande")
	default:
		fmt.Println("El número debe ser entero positivo.")
	}
}