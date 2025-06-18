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
# === goimports -local "github/CarolineMFVL/nls-auth" -w . ===

APP_NAME=nls-auth
PORT=4002

# === Commandes pour la base PostgreSQL ===

seed:
	@echo "Seeding database..."
	SEED_DB=1 go run main.go

reset:
	@echo "Resetting database..."
	go run tools/reset.go

# === Lancement de l'application Go ===

run:
	@echo "Lancement de  sur :"
	go run main.go

build:
	go build -o 

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
