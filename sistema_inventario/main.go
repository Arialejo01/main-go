package main

import (
	"fmt"
	"sistema_inventario/tienda"
)

func main() {
	inv := tienda.NuevoInventario()

	p1 := tienda.NuevoProducto(1, "Laptop Lenovo IdeaPad", 1200.50, 10)
	p2 := tienda.NuevoProducto(2, "Teclado Mecánico", 85.00, 25)

	inv.AgregarProducto(p1)
	inv.AgregarProducto(p2)

	fmt.Println("Inventario Inicial:")
	inv.ListarProductos()

	fmt.Println("\nActualizando precio y cantidad de la Laptop...")
	if p, existe := inv.BuscarProducto(1); existe {
		p.SetPrecio(1150.00)
		p.SetCantidad(8)
		inv.AgregarProducto(p) 
	}

	fmt.Println("\nEliminando el teclado del inventario...")
	inv.EliminarProducto(2)

	fmt.Println("\nInventario Final:")
	inv.ListarProductos()
}