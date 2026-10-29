package tienda

import "fmt"

type Producto struct {
	id       int
	nombre   string
	precio   float64
	cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		id:       id,
		nombre:   nombre,
		precio:   precio,
		cantidad: cantidad,
	}
}

func (p Producto) GetID() int { return p.id }
func (p Producto) GetNombre() string { return p.nombre }
func (p Producto) GetPrecio() float64 { return p.precio }
func (p Producto) GetCantidad() int { return p.cantidad }

func (p *Producto) SetPrecio(nuevoPrecio float64) { p.precio = nuevoPrecio }
func (p *Producto) SetCantidad(nuevaCantidad int) { p.cantidad = nuevaCantidad }

func (p Producto) MostrarInfo() {
	fmt.Printf("ID: %d | Nombre: %s | Precio: $%.2f | Cantidad: %d\n", p.id, p.nombre, p.precio, p.cantidad)
}