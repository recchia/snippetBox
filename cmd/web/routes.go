package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.LoadAndSave)
	router := http.NewServeMux()

	router.Handle("GET /snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippetForm))
	router.Handle("POST /snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippet))
	router.Handle("GET /snippet/{id}", dynamicMiddleware.ThenFunc(app.showSnippet))
	router.Handle("GET /", dynamicMiddleware.ThenFunc(app.home))

	router.Handle("GET /user/signup", dynamicMiddleware.ThenFunc(app.signUpUserForm))
	router.Handle("POST /user/signup", dynamicMiddleware.ThenFunc(app.signUpUser))
	router.Handle("GET /user/signin", dynamicMiddleware.ThenFunc(app.signInUserForm))
	router.Handle("POST /user/signin", dynamicMiddleware.ThenFunc(app.signInUser))
	router.Handle("POST /user/signout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.signOutUser))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(router)
}
