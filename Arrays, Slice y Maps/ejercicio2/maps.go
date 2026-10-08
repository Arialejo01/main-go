package main

import (
	"fmt"
	"strings"
)

func main() {
	// Map para almacenar actividades y votos
	votos := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	fmt.Println("=== Votación de Actividad de Integración ===")
	fmt.Println("Opciones: deportes, videojuegos, cine, musica")

	// Bucle para 5 votos válidos
	for i := 1; i <= 5; {
		fmt.Printf("Ingresa el voto %d: ", i)
		var entrada string
		fmt.Scan(&entrada)
		
		// Normalizar entrada a minúsculas
		entrada = strings.ToLower(strings.TrimSpace(entrada))

		// Verificar si la clave existe en el map
		if _, existe := votos[entrada]; existe {
			votos[entrada]++
			i++ // Solo avanza el contador si el voto es válido
		} else {
			fmt.Println("Actividad no válida. Intenta de nuevo.")
		}
	}

	fmt.Println("\n--- Resultados Finales ---")
	for actividad, cantidad := range votos {
		fmt.Printf("%s: %d votos\n", strings.ToUpper(actividad[0:1])+actividad[1:], cantidad)
	}

	ganador, maxVotos := determinarGanador(votos)
	fmt.Printf("\n La actividad ganadora es '%s' con %d votos.\n", ganador, maxVotos)
}

// Función que recibe el map y retorna la actividad con más votos
func determinarGanador(votos map[string]int) (string, int) {
	var actividadGanadora string
	maxVotos := -1

	for actividad, cantidad := range votos {
		if cantidad > maxVotos {
			maxVotos = cantidad
			actividadGanadora = actividad
		}
	}
	
	return actividadGanadora, maxVotos
}