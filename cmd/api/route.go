package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		cxt := context.WithValue(r.Context(), "params", p)
		next.ServeHTTP(w, r.WithContext(cxt))
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	secure := alice.New(app.checkToken)
	router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)
	router.HandlerFunc(http.MethodGet, "/test-ok", app.testOK)

	router.POST("/admin/delete-movie", app.wrap(secure.ThenFunc(app.deleteMovie)))
	router.POST("/create", app.wrap(secure.ThenFunc(app.createMenu)))

	return app.enableCORS(router)

}
