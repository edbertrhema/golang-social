package main

import "net/http"

func (app *application) HealthChecker(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))

}
