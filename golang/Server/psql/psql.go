package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


var db *sql.DB


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "buddhism"
)


func init() {
	var err error
	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully connected!")
}


func Close() {
	db.Close()
}


func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
