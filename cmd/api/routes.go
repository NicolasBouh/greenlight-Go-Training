package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/movies", app.ListMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.ShowMovieHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.CreateMovieHandler)
	router.HandlerFunc(http.MethodPut, "/v1/movies/:id", app.UpdateMovieHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.PatchMovieHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.DeleteMovieHandler)

	return app.recoverPanic(app.rateLimit(router))
}
