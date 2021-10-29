package data

import (
	"fmt"
	"log"
	"time"

	"github.com/paulndam/mock-chit-chat-forum/database"
)



type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}



// --------------- For Threads ------------------ --

func(user *User)  CreateAThread(topic string) (conv Thread ,err error){
	// var user *models.User

	statement := "insert into threads (uuid, topic, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, topic, user_id, created_at"

	stmt, err := database.DbConnection.SQL.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(CreateUUID(), topic, user.Id, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	
	return
	
}

func  Threads() (threads []Thread, err error){


	rows, err := database.DbConnection.SQL.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)

		fmt.Println("-------Thread ============>",threads)
		fmt.Println("----- All Threads from DB ------",rows)

	}
	fmt.Println("----- All Threads from DB ------",rows)
	rows.Close()

	return

}

// Gets a single thread by its UUID.
func  GetThreadByUUID(uuid string) (conv Thread, err error){

	conv = Thread{}

	query := "SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1"

	err = database.DbConnection.SQL.QueryRow(query,uuid).Scan(&conv.Id,&conv.Uuid,&conv.Topic,&conv.UserId,&conv.CreatedAt)


	return
}

// type Userx struct {

// 	User DBInterface
// }

// Gets user who started the thread.
func(thread *Thread)  User() (user User){

	user = User{}

	query := "SELECT id, uuid, name, email, created_at FROM users WHERE id = $1"

	database.DbConnection.SQL.QueryRow(query,thread.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)

	return
}


// ---------- POSTS ------------


// creates a post.
func(user *User)  CreatePost(conv Thread, body string) (post Post, err error){

	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"

	stmt,err := database.DbConnection.SQL.Prepare(statement)

	if err != nil {
		fmt.Println("---- Failed to insert Post data into to DB")
		return
	}

	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(CreateUUID(), body, user.Id, conv.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)

	return

}




// get posts to a thread
func(thread *Thread)  Posts()(posts []Post, err error){

	query := "SELECT id, uuid, body, user_id, thread_id, created_at FROM posts where thread_id = $1"

	rows, err := database.DbConnection.SQL.Query(query,thread.Id)

	if err != nil {
		fmt.Println("---- Inserting post to thread in DB failed")
		return
	}

	for rows.Next() {
		post := Post{}

		if err = rows.Scan(&post.Id,&post.Uuid,&post.Body,&post.UserId,&post.ThreadId,&post.CreatedAt); err != nil{
			fmt.Println("failed to scan row attributes for post into DB")
			return
		}
		posts = append(posts, post)
	}

	rows.Close()

	return
}

// Gets user who wrote the post.
func(post *Post)  User()(user User){

	user = User{}

	query := "SELECT id, uuid, name, email, created_at FROM users WHERE id = $1"

	database.DbConnection.SQL.QueryRow(query,post.UserId).Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.CreatedAt)

	return

}


// formating the createdAt data to display in a better date format on the screen.
func(thread *Thread)  CreatedAtDateThread() string{
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// formating the createdAt data to display in a better date format on the screen.
func(post *Post)  CreatedAtDatePost() string{
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}


// Gets number replies for of a post in a thread
func (thread *Thread)  NumReplies() (count int){

	

	query := `
		SELECT 
			count (*) 
		FROM 
			posts 
		where thread_id = $1
	`

	rows,err := database.DbConnection.SQL.Query(query,thread.Id)

	if err != nil {
		log.Println("error getting counts of threads from DB")
		return
	}

	for rows.Next(){
		if err = rows.Scan(&count); err != nil{
			return
		}
	}

	rows.Close()

	return
}