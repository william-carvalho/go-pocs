# tax-system

POC simples em Go para cadastro e consulta de regras de imposto e calculo de imposto por produto, estado e ano.

## Requisitos

- Go 1.22 ou superior

## Como executar

```bash
go run .
```

A API sobe em `http://localhost:8080`.

## Estrutura

```text
tax-system/
  main.go
  handlers/
  service/
  repository/
  model/
  dto/
```

## Endpoints

### POST /tax-rules

Cadastra uma nova regra de imposto.

Payload:

```json
{
  "product": "NOTEBOOK",
  "state": "SP",
  "year": 2024,
  "taxPercent": 0.12
}
```

### GET /tax-rules

Lista todas as regras cadastradas.

### GET /tax-rules/{product}/{state}/{year}

Busca uma regra especifica.

Exemplo:

```bash
curl http://localhost:8080/tax-rules/NOTEBOOK/SP/2024
```

### POST /tax/calculate

Calcula imposto com base na regra cadastrada.

Payload:

```json
{
  "product": "NOTEBOOK",
  "state": "SP",
  "year": 2024,
  "baseAmount": 1000.00
}
```

Resposta esperada:

```json
{
  "product": "NOTEBOOK",
  "state": "SP",
  "year": 2024,
  "baseAmount": 1000,
  "taxPercent": 0.12,
  "taxValue": 120,
  "totalAmount": 1120
}
```

## Exemplo rapido de teste

Cadastrar regra:

```bash
curl -X POST http://localhost:8080/tax-rules \
  -H "Content-Type: application/json" \
  -d "{\"product\":\"NOTEBOOK\",\"state\":\"SP\",\"year\":2024,\"taxPercent\":0.12}"
```

Calcular imposto:

```bash
curl -X POST http://localhost:8080/tax/calculate \
  -H "Content-Type: application/json" \
  -d "{\"product\":\"NOTEBOOK\",\"state\":\"SP\",\"year\":2024,\"baseAmount\":1000}"
```

## Regras e validacoes

- Nao permite regra duplicada para o mesmo produto + estado + ano
- Valida campos obrigatorios
- Nao permite `taxPercent` negativo
- Nao permite `baseAmount` negativo
- Retorna `404` quando a regra nao existe
- Retorna `409` para duplicidade
- Retorna `400` para payload invalido ou erro de validacao
