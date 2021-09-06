package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	r := mux.NewRouter()
	r.HandleFunc("/", app.home).Methods("GET")
	r.HandleFunc("/snippet/create", app.createSnippet).Methods("GET")
	r.HandleFunc("/snippet/create", app.createSnippet).Methods("POST")
	r.HandleFunc("/snippet/{id}", app.showSnippet).Methods("GET")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	return standardMiddleware.Then(r)
}
