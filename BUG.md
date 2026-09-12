# Defeito proposital

Este repositorio contem um defeito introduzido de proposito, usado nos testes de
ponta a ponta do aPRova. **Nao corrija.**

## Onde

- Arquivo: `go/frete/frete.go`
- Funcao: `frete.Calcular(pesoKg, distanciaKm float64) (float64, error)`

## Comportamento atual

Com `pesoKg` exatamente `0`, a funcao **retorna a tarifa base (`10.00`) e erro
`nil`**. A validacao de peso so rejeita valores estritamente negativos, entao o
zero atravessa as regras e cai na faixa "peso ate 1 kg".

```
frete.Calcular(0, 100)  // devolve 10.00, nil
```

## Comportamento esperado

Peso zero e entrada invalida e deve retornar erro (`ErrPesoInvalido`), sem valor
de frete — o mesmo tratamento ja dado ao peso negativo.

## Referencia correta

A implementacao **Python** (`python/frete/calculadora.py`) e a referencia: ela
levanta `PesoInvalidoError` para peso zero. A implementacao Java
(`java/src/main/java/com/aprova/frete/CalculadoraFrete.java`) concorda com a
Python. Somente a versao Go diverge.

## Detectabilidade

Nao existe no repositorio nenhum teste Go cobrindo peso zero. A suite Go passa
inteira (`go test ./...`) com o defeito presente: ele so aparece para quem
escrever o teste que falta.
