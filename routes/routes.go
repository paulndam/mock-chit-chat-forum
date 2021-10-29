package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/paulndam/mock-chit-chat-forum/auth_handler"
	"github.com/paulndam/mock-chit-chat-forum/home_handler"
	"github.com/paulndam/mock-chit-chat-forum/thread_handler"
)



func Routes() http.Handler{

	// creates new http handler
	mux := chi.NewRouter()
	// mux := http.NewServeMux()

	// middlewares.
	
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)

	mux.HandleFunc("/", home_handler.HomePage)
	mux.HandleFunc("/err", home_handler.ErrorPage)

	mux.HandleFunc("/login", auth_handler.LogIn)
	mux.HandleFunc("/logout", auth_handler.LogOut)
	mux.HandleFunc("/sign-up", auth_handler.SignUp)
	mux.HandleFunc("/signup-account", auth_handler.SignUpAccount)
	mux.HandleFunc("/authenticate", auth_handler.Authenticate)

	mux.HandleFunc("/thread/new", thread_handler.NewThread)
	mux.HandleFunc("/thread/create",thread_handler.CreateThread)
	mux.HandleFunc("/thread/all",thread_handler.GetAllThreadsFromDB)
	mux.HandleFunc("/thread/post",thread_handler.PostThread)
	mux.HandleFunc("/thread/read",thread_handler.ReadThread)



	mux.HandleFunc("/*",auth_handler.NotFound)








	

	return mux
	


}