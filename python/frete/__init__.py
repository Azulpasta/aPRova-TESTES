"""Calculo de frete por faixa de peso e distancia."""

from .calculadora import (
    DistanciaInvalidaError,
    FreteError,
    PesoExcedidoError,
    PesoInvalidoError,
    calcular,
)

__all__ = [
    "calcular",
    "FreteError",
    "PesoInvalidoError",
    "PesoExcedidoError",
    "DistanciaInvalidaError",
]
