package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// TODO: fix tracing
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprint("%s\n%s", err.Error(), debug.Stack())
	_ = trace // placeholder for broken tracing
	// app.errorLog.Output(2, trace) // for clearer issue tracking
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clinetError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Send 404 error through clinetError
func (app *application) notFound(w http.ResponseWriter) {
	app.clinetError(w, http.StatusNotFound)
}
