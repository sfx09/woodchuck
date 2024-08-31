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
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      req.Name,
	})

	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, "Unable to create new user"+err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (c *Controller) HandleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}

func (c *Controller) HandleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type Request struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	defer r.Body.Close()
	req := Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithErr(w, http.StatusBadRequest, "Unable to decode request JSON")
		return
	}
	feed, err := c.DB.CreateUserFeed(r.Context(), database.CreateUserFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      req.Name,
		Url:       req.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, "Failed to create new feed")
		return
	}
	respondWithJSON(w, http.StatusCreated, feed)
}
