package main

import (
	"errors" // Añadido para que compile errors.New
	"fmt"
	"time"
)

// 1
type Medicine struct {
	name             string
	manufacturer     string
	manufacturerDate time.Time
	shelfLife        int
}

// 2
func (m Medicine) GetName() string {
	return m.name
}

func (m *Medicine) SetName(nuevoNombre string) error {
	if nuevoNombre == "" {
		return errors.New("el nombre no puede estar vacío")
	}
	m.name = nuevoNombre
	return nil
}

func (m Medicine) GetManufacturer() string {
	return m.manufacturer
}

func (m *Medicine) SetManufacturer(nuevoManofactura string) error {
	if nuevoManofactura == "" {
		return errors.New("No puede estar vacío")
	}
	m.manufacturer = nuevoManofactura
	return nil
}

func (m Medicine) GetManufactureDate() time.Time {
	return m.manufacturerDate
}

func (m *Medicine) SetManufactureDate(nuevaFecha time.Time) {
	m.manufacturerDate = nuevaFecha
}

func (m Medicine) GetShelfLife() int {
	return m.shelfLife
}

func (m *Medicine) SetShelfLife(nuevaVidaUtile int) error {
	if nuevaVidaUtile <= 0 {
		return errors.New("la vida útil debe ser un número positivo")
	}
	m.shelfLife = nuevaVidaUtile
	return nil
}

// 3
type Tablet struct {
	Medicine
	dosePerTablet  int
	isPrescription bool
}

type Sytup struct {
	Medicine
	volume int
	flavor string
}

// 4
func NuevoTablet(medicine *Medicine, tableta int, receta bool) *Tablet {
	return &Tablet{
		Medicine:       *medicine,
		dosePerTablet:  tableta,
		isPrescription: receta,
	}
}

func NuevoSytup(medicine *Medicine, volumen int, sabor string) *Sytup {
	return &Sytup{
		Medicine: *medicine,
		volume:   volumen,
		flavor:   sabor,
	}
}

// 5
func (m *Tablet) MostrarInfoTablet() string {
	return fmt.Sprintf("Tableta: %d\n  Receta: %t\n",
		m.dosePerTablet,
		m.isPrescription,
	)
}

func (m *Sytup) MostrarInfoSytup() string {
	return fmt.Sprintf("Volumen: %d\n Sabor: %s\n",
		m.volume,
		m.flavor,
	)
}

// 6
func (m *Medicine) FechaCaducidad() time.Time {
	return m.manufacturerDate.AddDate(0, m.shelfLife, 0)
}

func main() {
	m1 := Medicine{
		name:             "Parcetamol",
		manufacturer:     "MedicinasEc",
		manufacturerDate: time.Now(),
		shelfLife:        24,
	}
	m2 := Medicine{
		name:             "sana sana",
		manufacturer:     "MedicinasEc",
		manufacturerDate: time.Now(),
		shelfLife:        12,
	}

	inventarioTabletas := []Tablet{
		*NuevoTablet(&m1, 500, false),
	}

	inventarioJarabes := []Sytup{
		*NuevoSytup(&m2, 120, "Vainilla"),
	}

	fmt.Println("Inv Tabletas")
	for _, t := range inventarioTabletas {
		fmt.Printf("Nombre: %s\n Empresa: %s\n", t.GetName(), t.GetManufacturer())
		fmt.Print(t.MostrarInfoTablet())
		fmt.Printf("Fecha Caducidad: %s\n\n", t.FechaCaducidad().Format("2026-12-22"))
	}

	fmt.Println("Inv Jarabes")
	for _, s := range inventarioJarabes {
		fmt.Printf("Nombre: %s\n Empresa: %s\n", s.GetName(), s.GetManufacturer())
		fmt.Print(s.MostrarInfoSytup())
		fmt.Printf("Fecha Caducidad: %s\n\n", s.FechaCaducidad().Format("2026-12-22"))
	}

	fmt.Println("****Actualisacion****")

	fmt.Printf("Nombre original del jarabe: %s\n", inventarioJarabes[0].GetName())

	inventarioJarabes[0].SetName("sana sana 2")

	fmt.Printf("Nombre modificado del jarabe: %s\n", inventarioJarabes[0].GetName())

}
