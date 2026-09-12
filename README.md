# aPRova — repositorio-alvo de testes

Este **nao e um sistema real**. E o repositorio usado como alvo pelo
[aPRova](https://github.com/), um GitHub App que intercepta Pull Requests,
executa os testes em ambiente isolado, classifica o risco da mudanca e condiciona
o merge. Aqui fica apenas o codigo sobre o qual o aPRova age.

O dominio e o mesmo nas tres linguagens — **calculo de frete por faixa de peso e
distancia** — com implementacoes independentes. Um dominio unico mantem o
benchmark comparavel entre linguagens: a diferenca observada vem do agente, nao
da complexidade do problema.

## Defeito proposital

A implementacao **Go** contem um defeito introduzido de proposito e documentado
em [BUG.md](BUG.md). Ele **nao deve ser corrigido**: e o caso de referencia dos
primeiros testes de ponta a ponta. A suite Go existente passa mesmo com o defeito
presente.

## Regras de negocio

| Condicao | Resultado |
| --- | --- |
| Peso ate 1 kg | Tarifa base de 10.00 |
| Peso acima de 1 kg | Tarifa base + 2.50 por quilo adicional, arredondado para cima no quilo |
| Distancia acima de 500 km | Acrescimo de 30% sobre o total |
| Peso acima de 50 kg | Erro: nao transportavel |
| Peso zero ou negativo | Erro: entrada invalida |
| Distancia negativa | Erro: entrada invalida |

## Rodando os testes

Cada pasta e um projeto completo e independente; os comandos rodam de dentro dela.

### Go

```bash
cd go
go test ./...
```

### Python

```bash
cd python
pip install -r requirements.txt
pytest
```

### Java

```bash
cd java
mvn test
```

## Estrutura

```
go/      pacote frete + frete_test.go
python/  pacote frete + tests/
java/    CalculadoraFrete + CalculadoraFreteTest (JUnit 5)
config/  settings.yaml
BUG.md   descricao do defeito proposital
```

## Arquivos de configuracao e CI

`.github/workflows/ci.yml`, `Dockerfile` e `config/settings.yaml` existem para
exercitar o **classificador de risco** do aPRova, que trata esses caminhos como
sensiveis. O conteudo e trivial e nao ha segredo nenhum. O workflow de CI,
porem, roda de verdade as tres suites em jobs separados, servindo tambem como
verificacao independente do repositorio.
