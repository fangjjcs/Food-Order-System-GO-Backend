package main

import (
	"errors"
	"net/http"
	"strconv"

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

// func (app *application) insertMovie(w http.ResponseWriter, r *http.Request) {
// }
// func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {
// }
// func (app *application) searchMovie(w http.ResponseWriter, r *http.Request) {
// }
