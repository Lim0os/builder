package domain

// Шаблоны файлов
const (
	ApplicationAppGoTemplate = `package application

type App struct {
	// TODO: добавьте зависимости приложения
}

func NewApp() *App {
	return &App{}
}`

	DomainEntityGoTemplate = `package domain

// TODO: определите ваши сущности
type Entity struct {
	ID string
}`

	DomainRepositoryGoTemplate = `package domain

type Repository interface {
	// TODO: определите методы репозитория
}`

	HttpServerGoTemplate = `package http_server

import "net/http"

type Server struct {
	// TODO: добавьте зависимости сервера
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, nil)
}`

	GoModTemplate = `module {{.}}

go 1.21

require (
	github.com/stretchr/testify v1.8.4
)`

	MainGoTemplate = `package main

import (
	"log/slog"
	"{{.}}/src/ports_adapters/primary/http_server"
)

func main() {
	server := http_server.NewServer()
	if err := server.Start(":8080"); err != nil {
		slog.Error("server failed", "error", err)
	}
}`

	MakefileTemplate = `.PHONY: build run test

build:
	go build -o bin/{{.}} main.go

run:
	go run main.go

test:
	go test ./...`

	DockerfileTemplate = `FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /{{.}}

EXPOSE 8080

CMD [ "/{{.}}" ]`

	ReadmeTemplate = `# {{.}}

Проект сгенерирован автоматически.`

	GitignoreTemplate = `bin/
vendor/
.config
*.log`
)
