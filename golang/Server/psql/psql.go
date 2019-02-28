package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func open() *sql.DB {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres "+
		"password=834295 dbname=postgres sslmode=disable")
	CheckErr(err)
	return db
}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

var db *sql.DB

func init() {
	db = open()
	fmt.Println("数据库连接成功")
}

func TestDB() {
	fmt.Print("数据库连接测试完成\n")
}
