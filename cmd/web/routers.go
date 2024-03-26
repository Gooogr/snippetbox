package main

import "net/http"

func (app *application) routers() *http.ServeMux {
	mux := http.NewServeMux()

	// TODO: turn off file listing
	// https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
