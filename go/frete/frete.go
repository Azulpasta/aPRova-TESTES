// Package frete calcula o valor do frete por faixa de peso e distância.
package frete

import (
	"errors"
	"math"
)

const (
	TarifaBase         = 10.00
	ValorPorQuiloExtra = 2.50
	PesoMaximoKg       = 50.0
	DistanciaLongaKm   = 500.0
	AcrescimoLongaDist = 0.30
)

var (
	// ErrPesoInvalido indica peso zero ou negativo.
	ErrPesoInvalido = errors.New("peso invalido: deve ser maior que zero")
	// ErrPesoExcedido indica carga acima do limite transportável.
	ErrPesoExcedido = errors.New("peso excedido: acima de 50 kg nao e transportavel")
	// ErrDistanciaInvalida indica distância negativa.
	ErrDistanciaInvalida = errors.New("distancia invalida: nao pode ser negativa")
)

// Calcular retorna o valor do frete para um peso em quilos e uma distância em
// quilômetros.
func Calcular(pesoKg, distanciaKm float64) (float64, error) {
	if pesoKg < 0 {
		return 0, ErrPesoInvalido
	}
	if pesoKg > PesoMaximoKg {
		return 0, ErrPesoExcedido
	}
	if distanciaKm < 0 {
		return 0, ErrDistanciaInvalida
	}

	total := TarifaBase
	if pesoKg > 1 {
		quilosExtras := math.Ceil(pesoKg - 1)
		total += quilosExtras * ValorPorQuiloExtra
	}
	if distanciaKm > DistanciaLongaKm {
		total *= 1 + AcrescimoLongaDist
	}
	return total, nil
}
