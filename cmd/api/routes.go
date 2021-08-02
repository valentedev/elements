package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.listElementsHandler)
	router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)

	return router
}
