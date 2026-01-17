# Go Messaging API

**nls-auth** is a backend microservice written in Go for authentication. It uses PostgreSQL as the database and GORM as the ORM.

## Features

- User authentication (registration and login)
- Data seeding to initialize the database
- REST API documented with Swagger

## Prerequisites

- [Go](https://golang.org/) 1.22 or higher
- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)
- PostgreSQL

## Swagger Documentation

Access at [http://localhost:4002/swagger/index.html]  
Or [http://localhost:4002/swagger/doc.json]

## Installations

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/go-messaging.git
   cd go-messaging
   ```

## Prepare API

@Todo: Verify and change variables if needed.

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

Start the application  
With Docker  
Start the services with Docker Compose:

`docker-compose up --build`

1. Create the role/user in PostgreSQL  
   In a terminal, run:

`psql -h localhost -p 5432 -U postgres -d nls_db`

postgres  
Then, in the psql shell:

````sql
CREATE ROLE cmf WITH LOGIN PASSWORD 'test1234';
ALTER ROLE cmf CREATEDB;
GRANT ALL ON SCHEMA public TO cmf;