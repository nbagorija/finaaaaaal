package main

import (
	"log"
	"os"
	"strconv"

	"github.com/nbagorija/finaaaaaal/server"
)

func main() {
	// Порт
	port := 7540

	// Проверка переменной окружения TODO_PORT
	if envPort := os.Getenv("TODO_PORT"); envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil {
			port = p
		} else {
			log.Printf("Неверное значение TODO_PORT %q, используется порт по умолчанию %d", envPort, port)
		}
	}

	// Запуск сервера
	if err := server.StartServer(port); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
