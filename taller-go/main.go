package main

import (
	"fmt"
	"strings"

	"github.com/Arialejo01/taller-go/utilidades"
)

func main() {
	var opcion string

	for {
		fmt.Println("\n=== Menú Principal ===")
		fmt.Println("1 Conversor de Monedas")
		fmt.Println("2 Contador de Vocales")
		fmt.Println("0 Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		if opcion == "0" || strings.ToLower(opcion) == "salir" {
			fmt.Println("Saliendo del programa...")
			break
		}

		switch opcion {
		case "1":
			utilidades.ConversorMonedas()
		case "2":
			utilidades.ContadorVocales()
		default:
			fmt.Println("Opción incorrecta, intente de nuevo.")
		}
	}
}