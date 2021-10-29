package home_handler

import (
	"net/http"

	"github.com/paulndam/mock-chit-chat-forum/data"
	"github.com/paulndam/mock-chit-chat-forum/utils"
)

func HomePage(w http.ResponseWriter, request *http.Request){

	threads,err := data.Threads()
	if err != nil {
		utils.ErrorMessage(w,request,"can't get all threads")
	}else{
		_,err := utils.CheckSession(w,request)

	if err != nil {
		utils.GenerateHTML(w,threads,"layout","public.navbar","index")
	}else{
		utils.GenerateHTML(w,threads,"layout","private.navbar","index")

	}
	}
	
}

func ErrorPage(w http.ResponseWriter, request *http.Request){
	vals := request.URL.Query()
	_, err := utils.CheckSession(w, request)
	if err != nil {
		utils.GenerateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		utils.GenerateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}