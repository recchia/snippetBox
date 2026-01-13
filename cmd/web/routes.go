package main

import (
	"net/http"

	"github.com/gorilla/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	router := pat.New()

	router.Get("/snippet/create", app.createSnippetForm)
	router.Post("/snippet/create", app.createSnippet)
	router.Get("/snippet/{id}", app.showSnippet)
	router.Get("/", app.home)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(router)
}
