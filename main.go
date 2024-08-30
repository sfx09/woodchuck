package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sfx09/woodchuck/controller"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	c := controller.NewController()

	mux := http.NewServeMux()
	srv := http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}
	mux.HandleFunc("GET /v1/healthz", c.HandleReadiness)
	mux.HandleFunc("GET /v1/err", c.HandleError)

	log.Println("Listening on addr:", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
