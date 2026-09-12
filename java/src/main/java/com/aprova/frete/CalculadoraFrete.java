package com.aprova.frete;

/** Calculo de frete por faixa de peso e distancia. */
public final class CalculadoraFrete {

    public static final double TARIFA_BASE = 10.00;
    public static final double VALOR_POR_QUILO_EXTRA = 2.50;
    public static final double PESO_MAXIMO_KG = 50.0;
    public static final double DISTANCIA_LONGA_KM = 500.0;
    public static final double ACRESCIMO_LONGA_DISTANCIA = 0.30;

    private CalculadoraFrete() {
    }

    /**
     * Retorna o valor do frete para um peso em quilos e uma distancia em quilometros.
     *
     * @throws IllegalArgumentException peso zero ou negativo, peso acima do limite
     *                                  transportavel, ou distancia negativa
     */
    public static double calcular(double pesoKg, double distanciaKm) {
        if (pesoKg <= 0) {
            throw new IllegalArgumentException("peso invalido: deve ser maior que zero");
        }
        if (pesoKg > PESO_MAXIMO_KG) {
            throw new IllegalArgumentException("peso excedido: acima de 50 kg nao e transportavel");
        }
        if (distanciaKm < 0) {
            throw new IllegalArgumentException("distancia invalida: nao pode ser negativa");
        }

        double total = TARIFA_BASE;
        if (pesoKg > 1) {
            total += Math.ceil(pesoKg - 1) * VALOR_POR_QUILO_EXTRA;
        }
        if (distanciaKm > DISTANCIA_LONGA_KM) {
            total *= 1 + ACRESCIMO_LONGA_DISTANCIA;
        }
        return total;
    }
}
