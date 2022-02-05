package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {

	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     VERSION,
	}

	err := app.writeJSON(w, http.StatusOK, currentStatus, "status")
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	app.logger.Println("id is", id)

	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) testOK(w http.ResponseWriter, r *http.Request) {

	type jsonResp struct {
		OK bool `json:"ok"`
	}

	res := jsonResp{
		OK: true,
	}

	err := app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {
	type jsonResp struct {
		OK bool `json:"ok"`
	}

	res := jsonResp{
		OK: true,
	}

	err := app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

type CreateParser struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Memo       string `json:"memo"`
	FileString string `json:"fileString"`
}

func (app *application) createMenu(w http.ResponseWriter, r *http.Request) {

	var parser CreateParser
	err := json.NewDecoder(r.Body).Decode(&parser)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	var request models.CreateRequest
	request.Name = parser.Name
	request.Type = parser.Type
	request.Memo = parser.Memo
	request.FileString = parser.FileString
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	err = app.models.DB.Create(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	type jsonResp struct {
		OK bool `json:"ok"`
	}

	res := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, res, "response")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}
