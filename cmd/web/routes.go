package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)
	router := http.NewServeMux()

	router.Handle("GET /snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	router.Handle("POST /snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet))
	router.Handle("GET /snippet/{id}", dynamicMiddleware.ThenFunc(app.showSnippet))
	router.Handle("GET /", dynamicMiddleware.ThenFunc(app.home))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(router)
}
