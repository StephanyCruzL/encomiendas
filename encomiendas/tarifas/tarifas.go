package tarifas

import (
	"errors"
	"fmt"
)

// TipoServicio representa la modalidad de envío.
type TipoServicio int

// TODO 1: declara con iota las constantes Estandar, Express y Prioritario.
const (
	Estandar TipoServicio = iota
	Express
	Prioritario
)

// TODO 2: declara las constantes TarifaBase, PrecioKg, PrecioKm, PesoMaximo e IGV.
const (
	TarifaBase = 8.00
	PrecioKg   = 2.50
	PrecioKm   = 0.15
	PesoMaximo = 50.00
	IGV        = 0.18
)

// CalcularSubtotal devuelve el costo sin IGV o un error si los datos no son válidos.
func CalcularSubtotal(pesoKg float64, distanciaKm int, servicio TipoServicio) (float64, error) {
	// TODO 3: valida el peso y la distancia; si no son válidos, devuelve 0 y un error.
	if pesoKg <= 0 || pesoKg > PesoMaximo {
		return 0, fmt.Errorf("peso no válido: %.2f kg (debe estar entre 0 y 50 kg)", pesoKg)
	}
	if distanciaKm <= 0 {
		return 0, errors.New("la distancia debe ser mayor a 0 km")
	}

	// TODO 4: obtén el recargo según el servicio y calcula el subtotal.
	recargo := 1.0
	switch servicio {
	case Express:
		recargo = 1.3
	case Prioritario:
		recargo = 1.6
	}

	subtotal := (TarifaBase + (pesoKg * PrecioKg) + (float64(distanciaKm) * PrecioKm)) * recargo
	return subtotal, nil
}

// AplicarIGV devuelve el IGV y el total usando retornos con nombre.
func AplicarIGV(subtotal float64) (igv, total float64) {
	// TODO 5: calcula igv y total.
	igv = subtotal * IGV
	total = subtotal + igv
	return
}