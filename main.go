package main

import (
	"net/http"

	"github.com/paulndam/mock-chit-chat-forum/routes"
)

const portNumber string = ":5000"


func main(){

	// create multiplexer that will redirect request to a handler
	mux := http.NewServeMux()
	// server all file at public path
	files := http.FileServer(http.Dir("/public"))
	// this will remove all strings starting with /static/ upon request and just render the file name it self.
	mux.Handle("/static/", http.StripPrefix("/static/",files))


	server := &http.Server{
		Addr: portNumber,
		Handler: routes.Routes(),
		
	}
	server.ListenAndServe()

}