package farmacia

import (
	"fmt"
	"time"
)

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int 
}

func (m *Medicine) SetName(nuevoNombre string) {
	if nuevoNombre != "" {
		m.name = nuevoNombre
	} else {
		fmt.Println("Error: El nombre no puede estar vacío")
	}
}

func (m *Medicine) GetName() string {
	return m.name
}

func (m *Medicine) GetManufactureDate() time.Time {
	return m.manufactureDate
}

func (m *Medicine) SetShelfLife(meses int) {
	if meses > 0 {
		m.shelfLife = meses
	}
}

// 6. Método para calcular caducidad (Suma los meses a la fecha de fabricación)
func (m *Medicine) CalcularCaducidad() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}

// ==========================================
// Estructura Tablet (Hija - usa composición)
// ==========================================
type Tablet struct {
	Medicine       
	dosePerTablet  float64
	isPrescription bool
}

func NuevoTablet(nombre, fabricante string, fecha time.Time, vidaUtil int, dosis float64, receta bool) Tablet {
	t := Tablet{}
	t.SetName(nombre) 
	t.manufacturer = fabricante
	t.manufactureDate = fecha
	t.SetShelfLife(vidaUtil)
	t.dosePerTablet = dosis
	t.isPrescription = receta
	return t
}

func (t *Tablet) SetDosis(nuevaDosis float64) {
	if nuevaDosis > 0 {
		t.dosePerTablet = nuevaDosis
	}
}

func (t *Tablet) GetDosis() float64 {
	return t.dosePerTablet
}

func (t *Tablet) MostrarDetalles() {
	fmt.Println("Tableta:", t.name, "| Fabricante:", t.manufacturer, "| Dosis:", t.dosePerTablet, "mg")
}

// ==========================================
// Estructura Syrup (Hija - usa composición)
// ==========================================
type Syrup struct {
	Medicine
	volume float64
	flavor string
}

func NuevoSyrup(nombre, fabricante string, fecha time.Time, vidaUtil int, volumen float64, sabor string) Syrup {
	s := Syrup{}
	s.SetName(nombre)
	s.manufacturer = fabricante
	s.manufactureDate = fecha
	s.SetShelfLife(vidaUtil)
	s.volume = volumen
	s.flavor = sabor
	return s
}

func (s *Syrup) MostrarDetalles() {
	fmt.Println("Jarabe:", s.name, "| Sabor:", s.flavor, "| Volumen:", s.volume, "ml")
}