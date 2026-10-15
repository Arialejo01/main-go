package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func calcularTotal(p Producto) float64 {
	return p.Precio * float64(p.Cantidad)
}

func estadoStock(p Producto) string {
	if p.Cantidad > 0 {
		return "Disponible"
	}
	return "Sin stock"
}

func mostrarProducto(p Producto) {
	fmt.Println("Nombre:     ", p.Nombre)
	fmt.Printf("Precio:      $%.2f\n", p.Precio)
	fmt.Println("Cantidad:   ", p.Cantidad)
	fmt.Printf("Valor total: $%.2f\n", calcularTotal(p))
	fmt.Println("Estado:     ", estadoStock(p))
	fmt.Println("------------------------------")
}

func main() {
	productos := []Producto{
		{Nombre: "Laptop", Precio: 800.50, Cantidad: 5},
		{Nombre: "Mouse", Precio: 15.00, Cantidad: 0},
		{Nombre: "Teclado", Precio: 30.00, Cantidad: 10},
	}

	fmt.Println("--- LISTA DE PRODUCTOS ---")
	for _, prod := range productos {
		mostrarProducto(prod)
	}
}