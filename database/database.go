package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)



type DB struct{
	SQL *sql.DB

}

var DbConnection = &DB{}


func ConnectToSQL(dsn string)(*DB, error){

	d,err := NewDatabase(dsn)

	if err != nil {
		panic(err)
	}

	DbConnection.SQL = d

	err = TestDB(d)
	if err != nil {
		return nil,err
	}

	return DbConnection,nil
}

// this func tries to ping the DB.
func TestDB(d *sql.DB) error{
	err := d.Ping()

	if err != nil {
		log.Println("Failed to Ping DB")
		return err
	}
	return nil
}


// creates a new DB.
func NewDatabase(dsn string)(*sql.DB,error){
	db,err := sql.Open("pgx",dsn)
	if err != nil {
		return nil,err
	}
	// check if were are getting a ping to DB and also if there's any error
	if err = db.Ping(); err != nil{
		return nil,err
	}

	return db,nil
}