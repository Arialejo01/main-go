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

func (m *Medicine) CalcularCaducidad() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}

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

