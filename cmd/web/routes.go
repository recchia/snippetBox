package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/recchia/snippetbox/ui"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.Handle("GET /static/", http.FileServerFS(ui.Files))

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	router.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	router.Handle("GET /about", dynamic.ThenFunc(app.about))
	router.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	router.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	router.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	router.Handle("GET /user/signin", dynamic.ThenFunc(app.userSignIn))
	router.Handle("POST /user/signin", dynamic.ThenFunc(app.userSignInPost))

	protected := dynamic.Append(app.requireAuthentication)

	router.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))
	router.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))
	router.Handle("POST /user/signout", protected.ThenFunc(app.userSignOutPost))
	router.Handle("GET /account/view", protected.ThenFunc(app.accountView))

	router.Handle("GET /ping", http.HandlerFunc(ping))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(router)
}
