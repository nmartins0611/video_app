package main

//GOOS=linux GOARCH=amd64 go build -o atease.exe

import (
	"net/http"
	"videoweb/web/root"

	"github.com/julienschmidt/httprouter"
)

var mux *httprouter.Router

//
// Initialize settings
func init() {
	// Setup the mux routing
	setupRouting()
}

//
// Set up the web server routing
func setupRouting() {
	mux = httprouter.New()

	mux.GET("/", root.Index)

	mux.ServeFiles("/assets/*filepath", http.Dir("./ui/static"))
}

//
// Main
func main() {
	doDevelopment()
}

//
// Listen on local port 10443
func doDevelopment() {
	http.ListenAndServeTLS(":10443", "keys/cert.pem", "keys/key.pem", mux)
}
