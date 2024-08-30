package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sfx09/woodchuck/internal/database"
)

type Controller struct {
	Port string
	DB   *database.Queries
}

func NewController(conn string) (Controller, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return Controller{}, err
	}
	dbQueries := database.New(db)
	return Controller{
		DB: dbQueries,
	}, nil
}

func (c *Controller) HandleReadiness(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, Response{Status: "ok"})
}

func (c *Controller) HandleError(w http.ResponseWriter, r *http.Request) {
	respondWithErr(w, http.StatusInternalServerError, "Internal server error")
}

func (c *Controller) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Name string `json:"name"`
	}
	req := Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithErr(w, http.StatusBadRequest, "Unable to decode request JSON")
		return
	}

	user, err := c.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      req.Name,
	})

	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, "Unable to create new user"+err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}
