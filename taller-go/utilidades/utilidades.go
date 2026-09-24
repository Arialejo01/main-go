package utilidades

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ConversorMonedas convierte dólares a diferentes divisas
func ConversorMonedas() {
	var dolares float64
	var opcionMoneda string

	fmt.Print("Ingrese el valor en dólares (USD): ")
	fmt.Scan(&dolares)

	fmt.Println("\nSeleccione la moneda a la que desea convertir:")
	fmt.Println("1 Euros")
	fmt.Println("2 LB (Libras Esterlinas)")
	fmt.Println("3 Won (Sur Coreano)")
	fmt.Println("4 BTC (Bitcoin)")
	fmt.Print("Opción: ")
	fmt.Scan(&opcionMoneda)

	tasaEuro := 0.92
	tasaLibra := 0.79
	tasaWon := 1340.50
	tasaBTC := 0.000015

	switch opcionMoneda {
	case "1":
		fmt.Printf("$%.2f USD equivalen a %.2f Euros\n", dolares, dolares*tasaEuro)
	case "2":
		fmt.Printf("$%.2f USD equivalen a %.2f Libras Esterlinas (LB)\n", dolares, dolares*tasaLibra)
	case "3":
		fmt.Printf("$%.2f USD equivalen a %.2f Won\n", dolares, dolares*tasaWon)
	case "4":
		fmt.Printf("$%.2f USD equivalen a %.6f BTC\n", dolares, dolares*tasaBTC)
	default:
		fmt.Println("Opción de moneda no válida.")
	}
}

// ContadorVocales cuenta las vocales de una frase ingresada
func ContadorVocales() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Ingrese una frase: ")
	scanner.Scan()
	frase := scanner.Text()
	if frase == "" {
		scanner.Scan()
		frase = scanner.Text()
	}

	frase = strings.ToLower(frase)

	a := strings.Count(frase, "a") + strings.Count(frase, "á")
	e := strings.Count(frase, "e") + strings.Count(frase, "é")
	i := strings.Count(frase, "i") + strings.Count(frase, "í")
	o := strings.Count(frase, "o") + strings.Count(frase, "ó")
	u := strings.Count(frase, "u") + strings.Count(frase, "ú")

	fmt.Println("\nResultados del conteo de vocales:")
	fmt.Printf("A: %d\nE: %d\nI: %d\nO: %d\nU: %d\n", a, e, i, o, u)
}