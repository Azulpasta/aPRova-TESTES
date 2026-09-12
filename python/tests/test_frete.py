import pytest

from frete import (
    DistanciaInvalidaError,
    PesoExcedidoError,
    PesoInvalidoError,
    calcular,
)


@pytest.mark.parametrize("peso", [0.1, 0.5, 1])
def test_peso_ate_um_quilo_usa_tarifa_base(peso):
    assert calcular(peso, 100) == pytest.approx(10.00)


def test_peso_acima_de_um_quilo_cobra_por_quilo_adicional():
    assert calcular(3, 100) == pytest.approx(15.00)


def test_quilo_adicional_arredonda_para_cima():
    assert calcular(2.1, 100) == pytest.approx(15.00)


def test_distancia_longa_aplica_acrescimo():
    assert calcular(1, 600) == pytest.approx(13.00)


def test_distancia_no_limite_nao_aplica_acrescimo():
    assert calcular(1, 500) == pytest.approx(10.00)


def test_peso_e_distancia_combinados():
    assert calcular(3, 600) == pytest.approx(19.50)


def test_peso_acima_do_limite_levanta_erro():
    with pytest.raises(PesoExcedidoError):
        calcular(50.1, 100)


def test_peso_no_limite_e_transportavel():
    assert calcular(50, 100) == pytest.approx(132.50)


def test_peso_zero_levanta_erro():
    with pytest.raises(PesoInvalidoError):
        calcular(0, 100)


def test_peso_negativo_levanta_erro():
    with pytest.raises(PesoInvalidoError):
        calcular(-1, 100)


def test_distancia_negativa_levanta_erro():
    with pytest.raises(DistanciaInvalidaError):
        calcular(1, -1)
