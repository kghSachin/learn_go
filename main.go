package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	fmt.Println("Hello, World!")
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT must be set")
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:  []string{"https://*", "http://*"},
		AllowedMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:  []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:  []string{"Link"},
		AllowCredentials: true,
		MaxAge:          300, // Maximum value not ignored by any of major browsers
	}))

	srv := &http.Server{
		Handler: router,
		Addr:   ":" + portString,
	}
	log.Printf("Server is running on port %s", portString)
	err:=srv.ListenAndServe()
	if err!=nil{
      log.Fatal(err)
	}

	fmt.Println("PORT is set to", portString)

}
