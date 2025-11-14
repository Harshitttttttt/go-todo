package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Harshitttttttt/go-todo/internal/config"
	"github.com/Harshitttttttt/go-todo/internal/database"
)

func main() {
	mux := http.NewServeMux()
	cfg := config.LoadConfig()

	// Connect Database
	db, err := database.InitDB(cfg.DatabaseUrl)
	if err != nil {
		log.Fatal("DB Connection Error")
	}
	defer db.Close()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	fmt.Printf("Environment Variables:\nPort: %d\nDatabase Url: %s\nEnvironment: %s\nJWT: %s\n", cfg.Port, cfg.DatabaseUrl, cfg.Environment, cfg.JWTSecret)
	fmt.Println("Server running at address: http://localhost:8080")
	http.ListenAndServe("localhost:8080", mux)
}
