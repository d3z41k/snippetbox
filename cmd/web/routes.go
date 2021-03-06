package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	r := mux.NewRouter()
	r.Handle("/", dynamicMiddleware.ThenFunc(app.home)).Methods("GET")
	r.Handle("/about", dynamicMiddleware.ThenFunc(app.about)).Methods("GET")

	r.Handle("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippetForm)).Methods("GET")
	r.Handle("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippet)).Methods("POST")
	r.Handle("/snippet/{id}", dynamicMiddleware.ThenFunc(app.showSnippet)).Methods("GET")

	r.Handle("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm)).Methods("GET")
	r.Handle("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser)).Methods("POST")
	r.Handle("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm)).Methods("GET")
	r.Handle("/user/login", dynamicMiddleware.ThenFunc(app.loginUser)).Methods("POST")
	r.Handle("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser)).Methods("POST")

	r.Handle("/user/profile", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.userProfile)).Methods("GET")
	r.Handle("/user/change-password", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.changePasswordForm)).Methods("GET")
	r.Handle("/user/change-password", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.changePassword)).Methods("POST")

	r.Handle("/ping", http.HandlerFunc(ping)).Methods("GET")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	return standardMiddleware.Then(r)
}
