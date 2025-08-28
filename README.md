# Truyenco

## Overview
This project contains a Go backend (Gin + GORM) and a React/TypeScript frontend built with Vite and TailwindCSS.

## Backend
- Gin REST API with JWT authentication and PostgreSQL via GORM.
- Environment variables configured via `.env`.
- Build and run:

```bash
go build ./cmd/server
./server
```

- Build binaries for Windows and Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o bin/server-linux ./cmd/server
GOOS=windows GOARCH=amd64 go build -o bin/server.exe ./cmd/server
```

- Lint with `golangci-lint run ./...`.

## Frontend
- React + TypeScript + Vite + TailwindCSS.
- Global state via Context API and `useReducer`.
- Forms handled with React Hook Form and Yup validation.
- Build frontend:

```bash
cd frontend
npm install
npm run build
```

- Lint with `npm run lint`.

## Docker
`docker-compose.yml` provides PostgreSQL, Redis and Kafka services for development.

## Configuration
Copy `.env.example` to `.env` and adjust values for database, Redis, Kafka and JWT secret.
