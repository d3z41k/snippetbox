package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	r := mux.NewRouter()
	r.Handle("/", dynamicMiddleware.ThenFunc(app.home)).Methods("GET")
	r.Handle("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm)).Methods("GET")
	r.Handle("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet)).Methods("POST")
	r.Handle("/snippet/{id}", dynamicMiddleware.ThenFunc(app.showSnippet)).Methods("GET")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	return standardMiddleware.Then(r)
}
