# MailFlow - Asynchronous Email Sending Service

A microservice for background email sending, built with Go, Redis (for queue), and SMTP.

## Features

- Asynchronous email sending using Redis queues
- Customizable HTML templates
- RESTful API for sending requests
- Docker containerization

## Requirements

- Go 1.16+
- Redis
- SMTP Server

## Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and configure the variables
3. Run `go mod download` to install dependencies

## Running the Service

### Locally

```bash
go run cmd/server/main.go