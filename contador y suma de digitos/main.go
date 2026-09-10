package main
import "fmt"

func main() {
	var n int
	fmt.Print("Ingresa un número entero positivo n: ")
	fmt.Scan(&n)

	temp := n
	contador := 0
	suma := 0

	if temp == 0 {
		contador = 1
		suma = 0
	} else {
		for temp > 0 {
			digito := temp % 10
			suma += digito
			contador++
			temp /= 10
		}
	}

	fmt.Printf("Cantidad de dígitos: %d\n", contador)
	fmt.Printf("Suma de los dígitos: %d\n", suma)

	switch contador {
	case 1:
		fmt.Println("Mensaje: Número de una cifra")
	case 2:
		fmt.Println("Mensaje: Número de dos cifras")
	case 3:
		fmt.Println("Mensaje: Número de tres cifras")
	default:
		if contador > 3 {
			fmt.Println("Mensaje: Número de varias cifras")
		} else {
			fmt.Println("Número inválido.")
		}
	}
}