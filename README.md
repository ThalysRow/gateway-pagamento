# Gateway de Pagamento

Este projeto foi desenvolvido durante a Imersão Full Cycle, onde construímos um Gateway de Pagamento completo utilizando arquitetura de microsserviços.
O objetivo é demonstrar a construção de um sistema distribuído moderno, com separação de responsabilidades, comunicação assíncrona e análise de fraudes em tempo real.

## Arquitetura

![image](https://github.com/user-attachments/assets/de14ab74-645d-453d-87a8-d0065803319e)


## Requisitos

- Go 1.24+
- Docker
- Docker Compose

## Configuração

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/gateway-pagamento.git
cd gateway-pagamento
```

2. Inicie o banco de dados:

```bash
docker compose up -d
```

3. Execute as migrações:

```bash
migrate -database "postgres://postgres:1234@localhost:5432/postgres?sslmode=disable" -path migrations up
```

4. Inicie a aplicação:

```bash
go run cmd/app/main.go
```

## Endpoints

### Contas

- POST /accounts - Criar uma nova conta
- GET /accounts - Obter dados da conta (requer X-API-Key)

### Faturas

- POST /invoice - Criar uma nova fatura (requer X-API-Key)

## Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_SSL_MODE=
HTTP_PORT=
```
