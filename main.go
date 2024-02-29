package main

import (
	"embed"
	handler "imaginai/handlers"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))

	port := os.Getenv("HTTP_LISTEN_ADDR")

	slog.Info("Server is running on port", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initEverything() error {
	return godotenv.Load()
}
