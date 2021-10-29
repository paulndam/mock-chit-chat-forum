package thread_handler

import (
	"fmt"
	"net/http"

	"github.com/paulndam/mock-chit-chat-forum/data"
	"github.com/paulndam/mock-chit-chat-forum/utils"
)

// ------------ Threads ----------

// Get
// Shows the new thread form page. check if user have session in order to make a new thread, if not redirect to login page.
func  NewThread(w http.ResponseWriter, request *http.Request){

	// get session by checking if user have session, thus calling the the checksession method
	_,err := utils.CheckSession(w,request)

	if err != nil {
		http.Redirect(w,request,"/login",http.StatusFound)
	}else{
		utils.GenerateHTML(w,request,"layout","private.navbar", "new.thread")
	}
}

// Post.
// creates a thread.
func  CreateThread(w http.ResponseWriter, request *http.Request){

	// get session by checking if user have session, thus calling the the checksession method
	sess,err := utils.CheckSession(w,request)

	if err != nil {
		http.Redirect(w,request,"/login",http.StatusFound)
	}else{
		err = request.ParseForm()

		if err != nil {
			utils.Danger("can't parse thread form", err)
		}

		user,err := sess.User()

		fmt.Println("---- User from session for CreateThread ----->",user)

		if err != nil {
			utils.Danger("can't get user from session",err)
		}

		topic := request.PostFormValue("topic")
		fmt.Println("--- Topic created by user is ------>",topic)

		

		if _,err := user.CreateAThread(topic); err != nil{
			fmt.Println("=========== Can't create Thread =======>",err)
		}


		http.Redirect(w,request,"/",http.StatusFound)
	}
	
}


// Get 
// Reads a thread, shows detail information of a thread, including the post and the form wrtie to post.
func  ReadThread(w http.ResponseWriter, request *http.Request){

	values := request.URL.Query()
	uuid := values.Get("id")
	thread,err := data.GetThreadByUUID(uuid)

	fmt.Println("-- Getting post of thread by its ID ----->",thread)

	if err != nil {
		utils.ErrorMessage(w,request,"----can't read thread------")
	}else{
		_,err := utils.CheckSession(w,request)
		
		if err != nil {
			utils.GenerateHTML(w,&thread,"layout","public.navbar","public.thread")
		}else{
			utils.GenerateHTML(w,&thread,"layout","private.navbar","private.thread")
		}
	}

}

// Get.
// Get all thread in DB.
func GetAllThreadsFromDB(w http.ResponseWriter, request *http.Request){
	threads,err := data.Threads()

	if err != nil {
		utils.ErrorMessage(w,request,"Can't get all threads.")
		return
	}else{
		_,err := utils.CheckSession(w,request)
		
		if err != nil {
			utils.GenerateHTML(w,&threads,"layout","public.navbar","index")
		}else{
			utils.GenerateHTML(w,&threads,"layout","private.navbar","index")
		}
	}
}


// Post.  
// Creates a post.
func  PostThread(w http.ResponseWriter, request *http.Request){

	sess,err := utils.CheckSession(w,request)

	if err != nil {
		http.Redirect(w,request,"/login",http.StatusFound)
	}else{
		err = request.ParseForm()

		if err != nil {
			utils.Danger("can't parse form",err)
		}

		user,err := sess.User()

		fmt.Println("---- User from session for CreateThread ----->",user)

		if err != nil {
			utils.Danger("can't get user from session",err)
		}

		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")

		thread, err := data.GetThreadByUUID(uuid)

		fmt.Println("---- post thread by id from thread_handler file ------->",thread)

		if err != nil {
			utils.ErrorMessage(w,request,"cannot read thread")
		}

		if _,err := user.CreatePost(thread,body); err != nil {
			utils.Danger("can't create post", err)
		}

		url := fmt.Sprint("/thread/read?id=",uuid)
		http.Redirect(w,request,url,http.StatusFound)
	}

}