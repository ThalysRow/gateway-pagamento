# Gateway de Pagamento

API de Gateway de Pagamento desenvolvida em Go.

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
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=1234
DB_NAME=postgres
DB_SSL_MODE=disable
HTTP_PORT=8080
```
