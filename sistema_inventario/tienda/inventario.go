package tienda

import "fmt"

type Inventario struct {
	productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		productos: make(map[int]Producto),
	}
}

func (i *Inventario) AgregarProducto(p Producto) {
	i.productos[p.GetID()] = p
}

func (i *Inventario) EliminarProducto(id int) {
	delete(i.productos, id)
}

func (i *Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := i.productos[id]
	return p, existe
}

func (i *Inventario) ListarProductos() {
	if len(i.productos) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}
	fmt.Println("--- Lista de Productos ---")
	for _, p := range i.productos {
		p.MostrarInfo()
	}
}