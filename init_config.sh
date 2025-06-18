#!/bin/bash

# filepath: /Users/carolinefauvel/Desktop/go/nls-go-auth/init_root_files.sh

set -e
# @todo: modify the variables below to match your project
APP_NAME="nls-auth"
MODULE_NAME="auth"
PORT=4002
DB_PORT=5432
DB_NAME="auth_db"

echo "Create initial files for $APP_NAME project..."

# .env
cat > .env <<EOF
PG_HOST="localhost"
PG_USER="postgres"
PG_PASSWORD="postgres"
PG_DB="$DB_NAME"
PG_PORT="$DB_PORT"
SEED_DB="0"
FAKE_USER="testuser"
FAKE_PASSWORD="testpassword"
FAKE_EMAIL="testuser@example.com"
EOF

# Dockerfile
cat > Dockerfile <<EOF
# Dockerfile
FROM golang:1.23.9

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend

EXPOSE $PORT
CMD ["./backend"]
EOF

# docker-compose.yml for database and adminer
cat > docker-compose.yml <<EOF
networks:
  postgres-network:
    driver: bridge

volumes:
  postgres-data:
    driver: local

services:
  # PostgreSQL
  messaging-postgres:
    image: postgres:16.9-bullseye
    container_name: postgres
    restart: unless-stopped
    ports:
      - "$PG_PORT:5432"
    volumes:
      - postgres-data:/var/lib/postgresql/data
    env_file:
      - .env
    environment:
      PGPASSWORD: $PG_PASSWORD
      POSTGRES_USER: $PG_USER
      POSTGRES_PASSWORD: $PG_PASSWORD
      POSTGRES_DB: $PG_DB
    networks:
      - postgres-network
  #adminer
  adminer:
    image: adminer:5.2.1
    container_name: adminer
    restart: unless-stopped
    ports:
      - "8080:8080"
    environment:
      ADMINER_DEFAULT_SERVER: postgres
    networks:
      - postgres-network

EOF

# README.md
cat > README.md <<EOF
# Go Messaging API

API de messagerie en Go avec Gorilla Mux, Swagger, PostgreSQL.

## Lancer en local

\`\`\`sh
docker-compose up
\`\`\`

## Documentation Swagger

Accédez à [http://localhost:4000/swagger/index.html](http://localhost:$DB_PORT/swagger/index.html)
EOF

# Makefile
cat > Makefile <<EOF
install: 
		go install github.com/air-verse/air@latest
		go install github.com/go-delve/delve/cmd/dlv@latest
		go install github.com/swaggo/swag/cmd/swag@latest
		go mod tidy

open-api:
		@echo "Génération de la documentation OpenAPI..."
		swag init --output ./docs --generalInfo main.go

format:
	@echo "🎨 Formatage du code Go..."
	go fmt ./...
# === goimports -local "github/CarolineMFVL/$APP_NAME" -w . ===

APP_NAME=$APP_NAME
PORT=$PORT

# === Commandes pour la base PostgreSQL ===

seed:
	@echo "Seeding database..."
	SEED_DB=1 go run main.go

reset:
	@echo "Resetting database..."
	go run tools/reset.go

# === Lancement de l'application Go ===

run:
	@echo "Lancement de $(APP_NAME) sur :$(PORT)"
	go run main.go

build:
	go build -o $(APP_NAME)

# === Tests ===

test-coverage:
	@echo "Exécution des tests avec couverture..."
	go test -coverprofile=coverage.out ./tests/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Couverture générée dans coverage.html"

test:
	@echo "Exécution des tests..."
	go test ./tests/...

# === Lint ===

lint:
	@echo "🔍 Linting du code Go..."
	golangci-lint run .

# === Docker ===

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down

docker-rebuild:
	docker-compose down && docker-compose up --build

# === Aide ===

help:
	@echo "Commandes disponibles :"
	@echo "  make run            - Launch app"
	@echo "  make seed           - Inject test users"
	@echo "  make reset          - Reset database"
	@echo "  make test           - Launch tests"
	@echo "  make build          - Build app"
	@echo "  make docker-up      - Launch docker"
	@echo "  make docker-down    - Stop docker"
	@echo "  make docker-rebuild - Rebuild complet Docker"
EOF

# .golangci.yml
cat > .golangci.yml <<EOF
linters:
  enable:
    - gofmt 
EOF

echo "Base files created successfully for $APP_NAME project."