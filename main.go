package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sfx09/woodchuck/controller"
)

func main() {
	godotenv.Load()

	dbConn := os.Getenv("CONN")
	c, err := controller.NewController(dbConn)

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	srv := http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}
	mux.HandleFunc("GET /v1/healthz", c.HandleReadiness)
	mux.HandleFunc("GET /v1/err", c.HandleError)
	mux.HandleFunc("POST /v1/users", c.HandleCreateUser)
	mux.HandleFunc("GET /v1/users", c.HandleAuthentication(c.HandleGetUser))
	mux.HandleFunc("POST /v1/feeds", c.HandleAuthentication(c.HandleCreateFeed))
	mux.HandleFunc("GET /v1/feeds", c.HandleGetFeeds)
	mux.HandleFunc("POST /v1/feed_follows", c.HandleAuthentication(c.HandleFollowFeed))
	mux.HandleFunc("GET /v1/feed_follows", c.HandleAuthentication(c.HandleGetFollows))
	mux.HandleFunc("DELETE /v1/feed_follows/{id}", c.HandleAuthentication(c.HandleDeleteFollow))

	go c.ScrapePeriodic()
	log.Println("Listening on addr:", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
