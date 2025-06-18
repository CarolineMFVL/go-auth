#!/bin/bash

# filepath: /Users/carolinefauvel/Desktop/go/init_api.sh

set -e

APP_NAME="nls-auth"
MODULE_NAME="nls-auth"
PORT=4001
ROOT=api/v1/auth

# Création de l'arborescence
mkdir -p $APP_NAME/{internal/handlers,internal/handlers/database,internal/models,internal/utils, internal/middlewares,tests,$ROOT,resources/public/docs}

cd $APP_NAME

# Initialisation du module Go
go mod init $MODULE_NAME

# Fichier .env exemple
cat > .env <<EOF
PG_HOST="localhost"
PG_USER="postgres"
PG_PASSWORD="postgres"
PG_DB="auth_db"
PG_PORT="5432"
EOF

# Swagger doc minimal
cat > resources/public/docs/open-api.yaml <<EOF
openapi: 3.0.0
info:
  title: API auth
  version: "1.0.0"
  description: API d'authentification exemple
servers:
  - url: http://localhost:$PORT
paths:
  /login:
    post:
      summary: Connexion utilisateur
      responses:
        '200':
          description: Connexion réussie
  /register:
    post:
      summary: Enregistrement utilisateur
      responses:
        '201':
          description: Utilisateur créé
EOF

echo "Structure du projet $APP_NAME créée avec succès !"