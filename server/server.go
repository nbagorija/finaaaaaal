package server

import (
	"fmt"
	"net/http"
	"os"
)

// StartServer запускает веб-сервер на указанном порту, обслуживая файлы из директории web.
func StartServer(port int) error {
	// Проверка существования директории web
	webDir := "./web"
	if _, err := os.Stat(webDir); os.IsNotExist(err) {
		return fmt.Errorf("директория web не существует: %v", err)
	}

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	// Запуск сервера
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Сервер запущен на http://localhost%s\n", addr)
	return http.ListenAndServe(addr, nil)
}
