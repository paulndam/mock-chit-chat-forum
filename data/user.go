package data

import (
	"fmt"
	"time"

	"github.com/paulndam/mock-chit-chat-forum/database"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}



func (user *User) CreateUser() (err error){

	// var user models.User
	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"

	stmt, err := database.DbConnection.SQL.Prepare(statement)

	if err != nil {
		
		return
	}

	
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow(CreateUUID(), user.Name, user.Email,Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return

}

// Get single user by UUID.
func  GetUserByUUID(uuid string)(user User, err error){

	user = User{}
	err = database.DbConnection.SQL.QueryRow("SELECT id,uuid,name,email,password,created_at FROM users WHERE uuid = $1", uuid).Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.Password,&user.CreatedAt)

	return
}

// Gets all users from DB.
func  GetAllUsers() (users []User ,err error){

	rows,err := database.DbConnection.SQL.Query("SELECT id,uuid,name,email,password,created_at FROM users ")

	if err != nil {
		return
	}

	for rows.Next() {
		user := User{}

		if err = rows.Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.Password,&user.CreatedAt);err != nil{
			return
		}

		users = append(users, user)

	}

	rows.Close()

	return
}

// Updates users.
func (user *User)  UpdateUser() (err error){

	// var user models.User

	statement := "update users set name = $2, email =$3 where id = $1"

	stmt,err := database.DbConnection.SQL.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_,err = stmt.Exec(user.Id,user.Name, user.Email)

	return
}


// Delete user from DB.
func (user *User) DeleteUser() (err error){

	// var user models.User

	statement := "delete from users where id = $1"

	stmt,err := database.DbConnection.SQL.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_,err = stmt.Exec(user.Id)

	return 
}

// deletes all users from DB.
func  DeleteAllUsers() (err error){
	statement := "delete from users"

	_,err = database.DbConnection.SQL.Exec(statement)

	if err != nil {
			return
	}

	return
}

// Gets user from session.
func(session *Session)  User() (user User,err error){
	// var session models.Session
	user = User{}
	err = database.DbConnection.SQL.QueryRow("SELECT id,uuid,name,email,created_at FROM users WHERE id = $1",session.UserId).Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.CreatedAt)

	return
}

// creates new session for existing user.
func (user *User)  CreateSession() (session Session, err error){

	// var user models.User

	statement := "insert into sessions(uuid,email,user_id,created_at) values ($1,$2,$3,$4) returning id, uuid,email, user_id, created_at "

	stmt,err := database.DbConnection.SQL.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	// query the row and scans the returned id into the session struct.
	err = stmt.QueryRow(CreateUUID(),user.Email,user.Id,time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId,&session.CreatedAt)

	return
} 

// gets session for existing user.
func(user *User) GetExistingUserSession() (session Session,err error){

	session = Session{}
	// var user models.User

	query := "SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1"

	err = database.DbConnection.SQL.QueryRow(query,user.Id).Scan(
		&session.Id, 
		&session.Uuid, 
		&session.Email, 
		&session.UserId,
		&session.CreatedAt,
	)

	return

}

// Deletes all session from DB.
func  SessionDeleteAll() (err error){
	statement := "delete from sessions"

	_,err = database.DbConnection.SQL.Exec(statement)

	return
}





// checks if session is valid in DB.
func (s *Session)  CheckSessionInDB() (valid bool, err error){

	// var session models.Session 

	query := "SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1"

	err = database.DbConnection.SQL.QueryRow(query,s.Uuid).Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId,&s.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if s.Id != 0 {
		valid = true
	}

	return
}


// gets user by email.

func  UserByEmail(email string) (user User, err error){

	user = User{}

	query := `

	SELECT 
		id,uuid,name,email,password,created_at 
	FROM 
		users 
	WHERE email = $1
	
	`

	err = database.DbConnection.SQL.QueryRow(query,email).Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.Password,&user.CreatedAt)

	return
}

// Delete's session fro DB.
func (session *Session)  DeleteByUUID() (err error){
	// var session models.Session
	statement := "delete from sessions where uuid = $1"
	stmt,err := database.DbConnection.SQL.Prepare(statement)
	if err != nil {
		fmt.Println("failed deleting session from DB",err)
		return
	}
	defer stmt.Close()

	_,err = stmt.Exec(session.Uuid)

	return
}
