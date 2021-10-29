package utils

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/paulndam/mock-chit-chat-forum/data"
	// "github.com/paulndam/chit-chat/models"
)

// function that will redirect to the error message page.
func ErrorMessage(w http.ResponseWriter, req *http.Request, msg string){
	url := []string{"/err?msg=",msg}
	http.Redirect(w,req,strings.Join(url,""),http.StatusFound)
}



// this utility function will check is a user is corrently login and has a session, or not
func   CheckSession(w http.ResponseWriter, req *http.Request) (sess data.Session, err error) {

	// get cookie from request
	cookie,err := req.Cookie("_cookie")

	// if no error and  there's session, that means user have  logged in, we retrieve the cookie stored 
	if err == nil{
		sess = data.Session{
			Uuid: cookie.Value,
		}
		
		// but if there's a session ,it checks the DB to see if the session Uuid exist

		if ok,_:= sess.CheckSessionInDB(); !ok {
			err = errors.New("invalid session")
		}
	}
	return

}

// parse in a list of filenames and gets a template
func ParseTemplateFiles(filenames ...string) (t *template.Template){
	var files []string 
	// initialize new template.
	t = template.New("layout")

	// loop thru each filenames and append a file.
	for _,file := range filenames {
		files = append(files,fmt.Sprintf("templates/%s.html",file))

	}
	t = template.Must(t.ParseFiles(files...))

	return


}

// will generate html templates.
// takes a response writer, some data and list of templates to be parsed
func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ...string){

	var files []string 
	for _,file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html",file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w,"layout",data)
}




// for logging purpose.

func Info(args ...interface{}){
	log.SetPrefix("INFO")
	log.Println(args...)
}

func Danger(args ...interface{}){
	log.SetPrefix("ERROR")
	log.Println(args...)
}

func Warning(args ...interface{}){
	log.SetPrefix("ERROR")
	log.Println(args...)
}

func Version() string{
	return "0.1"
}