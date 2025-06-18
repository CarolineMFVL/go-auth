#!/bin/bash

# filepath: /Users/carolinefauvel/Desktop/go/init_api.sh

set -e

APP_NAME="nls-auth"
MODULE_NAME="nls-auth"
PORT=4001
ROOT=api/v1/auth

# Create architecture directories
mkdir -p ./{constants/,docs/,internal/handlers,internal/handlers/database,internal/models,internal/utils, internal/middlewares,tests,$ROOT,resources/public/docs, api/v1/auth/db}

# Initialisation go
go mod init $MODULE_NAME

# Swagger doc minimal
cat > resources/public/docs/open-api.yaml <<EOF
openapi: 3.0.0
info:
  title: API auth
  version: "1.0.0"
  description: API authentication and authorization
servers:
  - url: http://localhost:$PORT
paths:
  /login:
    post:
      summary: User login
      responses:
        '200':
          description: Successful login
  /register:
    post:
      summary: Successful user registration
      responses:
        '201':
          description: User created successfully
EOF

echo " $APP_NAME global successfuly created !"