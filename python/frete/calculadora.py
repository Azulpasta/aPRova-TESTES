"""Regras de calculo de frete por faixa de peso e distancia."""

import math

TARIFA_BASE = 10.00
VALOR_POR_QUILO_EXTRA = 2.50
PESO_MAXIMO_KG = 50.0
DISTANCIA_LONGA_KM = 500.0
ACRESCIMO_LONGA_DISTANCIA = 0.30


class FreteError(ValueError):
    """Erro base do calculo de frete."""


class PesoInvalidoError(FreteError):
    """Peso zero ou negativo."""


class PesoExcedidoError(FreteError):
    """Peso acima do limite transportavel."""


class DistanciaInvalidaError(FreteError):
    """Distancia negativa."""


def calcular(peso_kg: float, distancia_km: float) -> float:
    """Retorna o valor do frete para um peso em quilos e uma distancia em km."""
    if peso_kg <= 0:
        raise PesoInvalidoError("peso invalido: deve ser maior que zero")
    if peso_kg > PESO_MAXIMO_KG:
        raise PesoExcedidoError("peso excedido: acima de 50 kg nao e transportavel")
    if distancia_km < 0:
        raise DistanciaInvalidaError("distancia invalida: nao pode ser negativa")

    total = TARIFA_BASE
    if peso_kg > 1:
        total += math.ceil(peso_kg - 1) * VALOR_POR_QUILO_EXTRA
    if distancia_km > DISTANCIA_LONGA_KM:
        total *= 1 + ACRESCIMO_LONGA_DISTANCIA
    return total
