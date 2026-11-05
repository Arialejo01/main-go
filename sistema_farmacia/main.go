package main

import (
	"fmt"
	"sistema_farmacia/farmacia"
	"time"
)

func main() {
	fecha1 := time.Date(2023, 5, 10, 0, 0, 0, 0, time.UTC)
	fechaHoy := time.Now()

	tabletas := []farmacia.Tablet{
		farmacia.NuevoTablet("Paracetamol", "Genfar", fecha1, 24, 500, false),
		farmacia.NuevoTablet("Amoxicilina", "MK", fechaHoy, 12, 875, true),
	}

	jarabes := []farmacia.Syrup{
		farmacia.NuevoSyrup("Ibuprofeno", "Bayer", fecha1, 36, 120, "Fresa"),
	}

	fmt.Println("--- INVENTARIO ---")
	for _, t := range tabletas {
		t.MostrarDetalles()
	}
	for _, j := range jarabes {
		j.MostrarDetalles()
	}

	fmt.Println("\n--- ACTUALIZACIÓN ---")
	fmt.Println("Dosis antigua:", tabletas[0].GetDosis(), "mg")
	
	tabletas[0].SetDosis(1000)
	
	fmt.Println("Dosis nueva (usando getter):", tabletas[0].GetDosis(), "mg")

	fmt.Println("\n--- CADUCIDAD ---")
	for _, t := range tabletas {
		fechaCaduca := t.CalcularCaducidad().Format("2006-01-02")
		fmt.Println(t.GetName(), "caduca el:", fechaCaduca)
	}

	for _, j := range jarabes {
		fechaCaduca := j.CalcularCaducidad().Format("2006-01-02")
		fmt.Println(j.GetName(), "caduca el:", fechaCaduca)
	}
}