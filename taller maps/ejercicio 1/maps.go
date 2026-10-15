package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Ingrese una frase o texto: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	palabras := strings.Fields(strings.ToLower(texto))

	conteo := make(map[string]int)
	for _, palabra := range palabras {
		conteo[palabra]++
	}

	fmt.Println("\n--- Conteo de palabras ---")
	palabraMasRepetida := ""
	maxRepeticiones := 0

	for palabra, cantidad := range conteo {
		fmt.Printf("%s: %d\n", palabra, cantidad)

		if cantidad > maxRepeticiones {
			maxRepeticiones = cantidad
			palabraMasRepetida = palabra
		}
	}

	fmt.Printf("\nLa palabra que más se repite es '%s' (%d veces).\n", palabraMasRepetida, maxRepeticiones)
}