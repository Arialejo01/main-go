
package main

import "fmt"

func main() {
	var opcion string

	for {
		fmt.Println("\n--- MENÚ PRINCIPAL ---")
		fmt.Println("1. Calcular promedio de estudiantes")
		fmt.Println("2. Calcular suma del 1 al N")
		fmt.Println("3. Convertir Celsius a Fahrenheit")
		fmt.Println("4. Convertir Fahrenheit a Celsius")
		fmt.Println("0 o 'salir' para terminar")
		fmt.Print("Elige una opción: ")
		fmt.Scan(&opcion)

		if opcion == "0" || opcion == "salir" {
			fmt.Println("Saliendo del programa...")
			break
		}

		switch opcion {
		case "1":
			opcion1()
		case "2":
			opcion2()
		case "3":
			opcion3()
		case "4":
			opcion4()
		default:
			fmt.Println("Opción no válida. Inténtalo de nuevo.")
		}
	}
}

// Función con parámetros y retorno para calcular el promedio
func averageGrade(sumaTotal float64, cantidad int) float64 {
	return sumaTotal / float64(cantidad)
}

// Funciones que no retornan valores para cada opción del menú
func opcion1() {
	var n int
	fmt.Print("Ingresa la cantidad de estudiantes: ")
	fmt.Scan(&n)

	var sumaNotas float64
	for i := 1; i <= n; i++ {
		var nota float64
		fmt.Print("Ingresa la nota (0 a 100) del estudiante ", i, ": ")
		fmt.Scan(&nota)
		// Vamos acumulando la nota sin necesidad de listas complejas
		sumaNotas = sumaNotas + nota 
	}

	// Llamamos a la función para que nos devuelva el cálculo
	promedio := averageGrade(sumaNotas, n)
	fmt.Println("\nEl promedio del curso es:", promedio)

	if promedio >= 70 {
		fmt.Println("Estado: APROBADO")
	} else {
		fmt.Println("Estado: REPROBADO")
	}

	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Rendimiento: Excellent performance")
	case promedio >= 80 && promedio <= 89:
		fmt.Println("Rendimiento: Good performance")
	case promedio >= 70 && promedio <= 79:
		fmt.Println("Rendimiento: Satisfactory performance")
	case promedio < 70:
		fmt.Println("Rendimiento: Needs improvement")
	}
}

func opcion2() {
	var n int
	fmt.Print("Ingresa un número entero N: ")
	fmt.Scan(&n)

	var suma int
	for i := 1; i <= n; i++ {
		suma = suma + i
	}
	fmt.Println("La suma del 1 al", n, "es:", suma)
}

func opcion3() {
	var celsius float64
	fmt.Print("Ingresa la temperatura en Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Println("La temperatura en Fahrenheit es:", fahrenheit)
}

func opcion4() {
	var fahrenheit float64
	fmt.Print("Ingresa la temperatura en Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println("La temperatura en Celsius es:", celsius)
}