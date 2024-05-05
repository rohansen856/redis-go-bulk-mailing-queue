FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar código-fonte
COPY . .

# Compilar aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o mailflow ./cmd/server

# Imagem final
FROM alpine:3.17

WORKDIR /app

# Copiar binário compilado
COPY --from=builder /app/mailflow .

# Criar diretório para templates
COPY --from=builder /app/internal/templates/html ./internal/templates/html

# Expor porta
EXPOSE 8080

# Comando para iniciar a aplicação
CMD ["./mailflow"]