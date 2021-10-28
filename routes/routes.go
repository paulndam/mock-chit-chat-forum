package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)



func Routes() http.Handler{

	// creates new http handler
	mux := chi.NewRouter()
	// mux := http.NewServeMux()

	// middlewares.
	
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)

	

	return mux
	


}