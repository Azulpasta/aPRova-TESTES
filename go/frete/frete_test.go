package frete

import (
	"errors"
	"math"
	"testing"
)

func quaseIgual(a, b float64) bool {
	return math.Abs(a-b) < 0.0001
}

func TestPesoAteUmQuiloUsaTarifaBase(t *testing.T) {
	casos := []float64{0.1, 0.5, 1}
	for _, peso := range casos {
		valor, err := Calcular(peso, 100)
		if err != nil {
			t.Fatalf("peso %v: erro inesperado: %v", peso, err)
		}
		if !quaseIgual(valor, 10.00) {
			t.Errorf("peso %v: esperado 10.00, obtido %v", peso, valor)
		}
	}
}

func TestPesoAcimaDeUmQuiloCobraPorQuiloAdicional(t *testing.T) {
	valor, err := Calcular(3, 100)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if !quaseIgual(valor, 15.00) {
		t.Errorf("esperado 15.00, obtido %v", valor)
	}
}

func TestQuiloAdicionalArredondaParaCima(t *testing.T) {
	valor, err := Calcular(2.1, 100)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if !quaseIgual(valor, 15.00) {
		t.Errorf("esperado 15.00, obtido %v", valor)
	}
}

func TestDistanciaLongaAplicaAcrescimo(t *testing.T) {
	valor, err := Calcular(1, 600)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if !quaseIgual(valor, 13.00) {
		t.Errorf("esperado 13.00, obtido %v", valor)
	}
}

func TestDistanciaNoLimiteNaoAplicaAcrescimo(t *testing.T) {
	valor, err := Calcular(1, 500)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if !quaseIgual(valor, 10.00) {
		t.Errorf("esperado 10.00, obtido %v", valor)
	}
}

func TestPesoEDistanciaCombinados(t *testing.T) {
	valor, err := Calcular(3, 600)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if !quaseIgual(valor, 19.50) {
		t.Errorf("esperado 19.50, obtido %v", valor)
	}
}

func TestPesoAcimaDoLimiteRetornaErro(t *testing.T) {
	if _, err := Calcular(50.1, 100); !errors.Is(err, ErrPesoExcedido) {
		t.Errorf("esperado ErrPesoExcedido, obtido %v", err)
	}
}

func TestPesoNoLimiteEhTransportavel(t *testing.T) {
	valor, err := Calcular(50, 100)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if !quaseIgual(valor, 132.50) {
		t.Errorf("esperado 132.50, obtido %v", valor)
	}
}

func TestPesoNegativoRetornaErro(t *testing.T) {
	if _, err := Calcular(-1, 100); !errors.Is(err, ErrPesoInvalido) {
		t.Errorf("esperado ErrPesoInvalido, obtido %v", err)
	}
}

func TestDistanciaNegativaRetornaErro(t *testing.T) {
	if _, err := Calcular(1, -1); !errors.Is(err, ErrDistanciaInvalida) {
		t.Errorf("esperado ErrDistanciaInvalida, obtido %v", err)
	}
}
