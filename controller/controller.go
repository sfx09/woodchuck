package controller

import "net/http"

type Controller struct {
	Port string
}

func NewController() Controller {
	return Controller{}
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
