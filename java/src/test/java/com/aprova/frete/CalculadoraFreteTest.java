package com.aprova.frete;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.ValueSource;

class CalculadoraFreteTest {

    private static final double TOLERANCIA = 0.0001;

    @ParameterizedTest
    @ValueSource(doubles = {0.1, 0.5, 1})
    @DisplayName("peso ate 1 kg usa a tarifa base")
    void pesoAteUmQuiloUsaTarifaBase(double peso) {
        assertEquals(10.00, CalculadoraFrete.calcular(peso, 100), TOLERANCIA);
    }

    @Test
    @DisplayName("peso acima de 1 kg cobra por quilo adicional")
    void pesoAcimaDeUmQuiloCobraPorQuiloAdicional() {
        assertEquals(15.00, CalculadoraFrete.calcular(3, 100), TOLERANCIA);
    }

    @Test
    @DisplayName("quilo adicional arredonda para cima")
    void quiloAdicionalArredondaParaCima() {
        assertEquals(15.00, CalculadoraFrete.calcular(2.1, 100), TOLERANCIA);
    }

    @Test
    @DisplayName("distancia acima de 500 km aplica acrescimo de 30%")
    void distanciaLongaAplicaAcrescimo() {
        assertEquals(13.00, CalculadoraFrete.calcular(1, 600), TOLERANCIA);
    }

    @Test
    @DisplayName("distancia exatamente 500 km nao aplica acrescimo")
    void distanciaNoLimiteNaoAplicaAcrescimo() {
        assertEquals(10.00, CalculadoraFrete.calcular(1, 500), TOLERANCIA);
    }

    @Test
    @DisplayName("peso e distancia combinados")
    void pesoEDistanciaCombinados() {
        assertEquals(19.50, CalculadoraFrete.calcular(3, 600), TOLERANCIA);
    }

    @Test
    @DisplayName("peso acima de 50 kg nao e transportavel")
    void pesoAcimaDoLimiteLancaErro() {
        assertThrows(IllegalArgumentException.class, () -> CalculadoraFrete.calcular(50.1, 100));
    }

    @Test
    @DisplayName("peso de exatamente 50 kg e transportavel")
    void pesoNoLimiteEhTransportavel() {
        assertEquals(132.50, CalculadoraFrete.calcular(50, 100), TOLERANCIA);
    }

    @Test
    @DisplayName("peso zero e entrada invalida")
    void pesoZeroLancaErro() {
        assertThrows(IllegalArgumentException.class, () -> CalculadoraFrete.calcular(0, 100));
    }

    @Test
    @DisplayName("peso negativo e entrada invalida")
    void pesoNegativoLancaErro() {
        assertThrows(IllegalArgumentException.class, () -> CalculadoraFrete.calcular(-1, 100));
    }

    @Test
    @DisplayName("distancia negativa e entrada invalida")
    void distanciaNegativaLancaErro() {
        assertThrows(IllegalArgumentException.class, () -> CalculadoraFrete.calcular(1, -1));
    }
}
