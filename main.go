package main

import (
	"log"
	"net/http"

	"github.com/paulndam/mock-chit-chat-forum/database"
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

	db,err := run()

	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

	server := &http.Server{
		Addr: portNumber,
		Handler: routes.Routes(),
		
	}
	server.ListenAndServe()

}

func run()(*database.DB,error){
	db,err := database.ConnectToSQL("host=localhost port=5432 dbname=mock-chit-chat") 
	if err != nil {
		log.Fatal("---- Connection to DataBase Failed---- ")
	}
	log.Println("----Connection to DataBase Established successfully --------")
	
	return db,nil
}