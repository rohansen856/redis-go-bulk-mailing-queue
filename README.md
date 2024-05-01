# MailFlow - Serviço de Envio de E-mails Assíncrono

Um microserviço para envio de e-mails em segundo plano, construído com Go, Redis (para fila) e SMTP.

## Características

- Envio assíncrono de e-mails usando filas Redis
- Templates HTML personalizáveis
- API RESTful para solicitações de envio
- Containerização com Docker

## Requisitos

- Go 1.16+
- Redis
- Servidor SMTP (ou serviço como SendGrid, Mailgun, etc.)

## Configuração

1. Clone o repositório
2. Copie `.env.example` para `.env` e configure as variáveis
3. Execute `go mod download` para instalar as dependências

## Executando o Serviço

### Localmente

```bash
go run cmd/server/main.go