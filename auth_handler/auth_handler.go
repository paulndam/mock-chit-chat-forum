package auth_handler

import (
	"fmt"
	"net/http"

	"github.com/paulndam/mock-chit-chat-forum/data"
	"github.com/paulndam/mock-chit-chat-forum/utils"
)

// renders login page.
func  LogIn(w http.ResponseWriter,request *http.Request){
	// calls the ParseTemplateFiles func which parses the files and renders it
	loginPage := utils.ParseTemplateFiles("login.layout","public.navbar","login")
	loginPage.Execute(w,nil)
}

// renders sign up page.
 func  SignUp(w http.ResponseWriter, request *http.Request){
	 utils.GenerateHTML(w,nil,"login.layout","public.navbar","signup")
 }

//  renders a 404, page not found.
func NotFound(w http.ResponseWriter, request *http.Request){
	utils.GenerateHTML(w,request,"layout","public.navbar","pageNotFound")
}

//  creates a user account.
func  SignUpAccount(w http.ResponseWriter, request *http.Request){
	// parse the form.
	err := request.ParseForm()

	if err != nil {
		utils.Danger(err, "can't parse form")
	}

	user := data.User{}
	user.Name = request.PostFormValue("name")
	user.Email = request.PostFormValue("email")
	user.Password = request.PostFormValue("password")

	if err := user.CreateUser();err != nil {
			// utils.Danger("error creating user from create user func",err)
			fmt.Println("error creating user",err)
	}

	fmt.Println("user created ---------->",user)

	// getting form object in order to check for form validation.

	// ------- TODO LATER ON ------
	// form := formvalidation.New(request.Form)

	// form.Required("name","email","password")
	// form.MinLen("name",2)
	// form.IsEmailValid("email")

	
	// // if form is not valid throw error.
	// if !form.Valid(){
	// 	data := make(map[string]interface{})
	// 	data["user"] = user

	// 	utils.GenerateHTML(w,request,"layout","public.navbar","signup")
	// }

	http.Redirect(w,request,"/login",http.StatusFound)


}

// Post request for authenticate.
// authenticates user by email and password.
func Authenticate(w http.ResponseWriter, request *http.Request){
	// parse the form.
	err := request.ParseForm()

	// get user by email and return if there's an error.
	user,err := data.UserByEmail(request.PostFormValue("email"))

	if err != nil {
		fmt.Println("user with email not found")
		utils.Danger("can't find user with that email ",err)
	}
	// check for password.
	if user.Password == data.Encrypt(request.PostFormValue("password")){
		// if password matches , create a session and store it.
		session,err := user.CreateSession()
		if err != nil{
			fmt.Println("can't create session for user")
			utils.Danger("can't create session for user",err)
		}
		fmt.Println("----Session created for user and now setting up cookie. ------")
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly: true,
		}
		// setting the cookie to response header.
		http.SetCookie(w,&cookie)
		http.Redirect(w,request,"/",http.StatusFound)
	}else{
		// if password doesn't match then redirect them login.
		http.Redirect(w,request,"/login",http.StatusFound)
	}

}

// Get method for logout.
// logsout user.
func LogOut(w http.ResponseWriter, request *http.Request){

	// get cookie from session and potential error.
	cookie,err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		utils.Warning("Failed to get cookie and user logging out failed",err)
		session := data.Session{
			Uuid: cookie.Value,
		}
		session.DeleteByUUID()
	}
	http.Redirect(w,request,"/",http.StatusFound)
	

}
