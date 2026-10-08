package main

import (
	"fmt"
)

func main() {
	// Array bidimensional de 6 estudiantes x 4 materias
	notas := [6][4]float64{
		{8.5, 9.0, 7.5, 8.0},
		{6.0, 7.5, 8.0, 9.5},
		{10.0, 9.5, 9.0, 10.0},
		{5.5, 6.0, 6.5, 7.0},
		{9.0, 8.5, 9.5, 9.0},
		{7.5, 8.0, 7.0, 7.5},
	}

	var sumaClase float64
	var totalNotas int

	fmt.Println("=== Análisis de Notas ===")

	for i := 0; i < len(notas); i++ {
		// Se crea un slice a partir de la fila del array
		sliceEstudiante := notas[i][:]
		
		promedio := calcularPromedio(sliceEstudiante)
		notaMax, notaMin := obtenerMaxMin(sliceEstudiante)

		fmt.Printf("Estudiante %d | Promedio: %.2f | Max: %.2f | Min: %.2f\n", i+1, promedio, notaMax, notaMin)

		// Acumular para el promedio general
		for _, nota := range sliceEstudiante {
			sumaClase += nota
			totalNotas++
		}
	}

	promedioGeneral := sumaClase / float64(totalNotas)
	fmt.Printf("\nPromedio general de la clase: %.2f\n", promedioGeneral)
}

// Recibe un slice de float64
func calcularPromedio(notas []float64) float64 {
	var suma float64
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}

// Recibe un slice y retorna dos valores float64
func obtenerMaxMin(notas []float64) (float64, float64) {
	max := notas[0]
	min := notas[0]
	
	for _, nota := range notas {
		if nota > max {
			max = nota
		}
		if nota < min {
			min = nota
		}
	}
	return max, min
}