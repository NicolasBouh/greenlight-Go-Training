package main

import (
	"greenlight.nicolasbouhours.fr/internal/data"
	"greenlight.nicolasbouhours.fr/internal/validator"
	"net/http"
	"time"
)

func (app *application) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJson(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	v := validator.New()
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	app.writeJson(w, http.StatusCreated, envelope{"movie": input}, nil)
}

func (app *application) ShowMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Year:      1972,
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJson(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
