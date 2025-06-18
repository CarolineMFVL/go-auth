# Go Messaging API

**nls-auth** est un microservice backend écrite en Go permettant une authentification. Elle utilise PostgreSQL comme base de données et GORM comme ORM.

## Fonctionnalités

- Authentification des utilisateurs (inscription et connexion)
- Seed de données pour initialiser la base
- API REST documentée avec Swagger

## Prérequis

- [Go](https://golang.org/) 1.22 ou supérieur
- [Docker](https://www.docker.com/) et [Docker Compose](https://docs.docker.com/compose/)
- PostgreSQL

## Documentation Swagger

Accédez à [http://localhost:4000/swagger/index.html](http://localhost:5433/swagger/index.html)

## Installation

1. Clonez le repository :

   ```bash
   git clone https://github.com/votre-utilisateur/go-messaging.git
   cd go-messaging
   ```

## Prepare API

@Todo : Verify and change variables if needed.

Run

`bash init_api.sh`

Then

`bash init_config.sh`

Then

`bash init_based_files.sh`

Then

`make install`

Then

`make open-api`

Lancer l'application
Avec Docker
Démarrez les services avec Docker Compose :

`docker-compose up --build`

1. Créez le rôle/utilisateur dans PostgreSQL
   Dans un terminal, lancez :

`psql -h localhost -p 5433 -U postgres -d nls_db`

postgres
Puis, dans le shell psql :

````CREATE ROLE cmf WITH LOGIN PASSWORD 'test1234';
ALTER ROLE cmf CREATEDB;
GRANT ALL ON SCHEMA public TO cmf;```

L'application sera disponible sur http://localhost:4000.

En local
Lancez PostgreSQL et configurez les variables d'environnement nécessaires (PG_HOST, PG_USER, PG_PASSWORD, PG_DB, PG_PORT).

Lancez l'application :

`go run main.go`

Documentation API
La documentation Swagger est générée automatiquement. Pour la générer, utilisez la commande suivante :

`make open-api`

Tests
Pour exécuter les tests, utilisez la commande suivante :

`go test ./...`


````
