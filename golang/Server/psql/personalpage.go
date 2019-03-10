package psql

import (
	// "database/sql"
	"fmt"
	"log"
	//_ "github.com/bmizerany/pq"
)

type User struct {
	UserId       int
	UserName     string
	Password     string
	Email        string
	Image        string
	Integral     int
	RegisterDate string
}

var user1 User
var err error

func SelectUser() User {
	rows, err := db.Query("SELECT * FROM users") //执行一次查询，返回多行结果
	if err != nil {
		log.Fatal(err)
	}
	//如果Next()返回假，rows回自动close()
	for rows.Next() {
		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。
		if err := rows.Scan(&user1.UserId, &user1.UserName, &user1.Password, &user1.Email, &user1.Image, &user1.Integral, &user1.RegisterDate); err != nil {
			log.Fatal(err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	//rows.Close()
	fmt.Println("数据查询成功！")
	fmt.Println(user1)
	return user1
}
func UpdateUser(user User) {
	//大写会自动转换为小写,要加斜杆和双引号
	stmt, err := db.Prepare("UPDATE \"users\" SET \"username\"=$1, \"password\"=$2, \"email\"=$3, \"image\"=$4, \"integral\"=$5, \"registrationdate\"=$6  WHERE \"userid\"=$7")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(user.UserName, user.Password, user.Email, user.Image, user.Integral, user.RegisterDate, user.UserId)
	stmt.Close()
	fmt.Println("修改成功！")
}
