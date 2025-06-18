# Dockerfile
FROM golang:1.23.9

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend

EXPOSE 4002
CMD ["./backend"]
