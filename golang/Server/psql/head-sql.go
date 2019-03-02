package psql

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// var db *sql.DB

// func Connectdb() {
// 	var err error
// 	db, err = sql.Open("postgres", "user=postgres password=zzm19981105 dbname=user sslmode=disable")
// 	//defer db.Close()
// 	check(err)
// }

func CheckEmailPassword(o Output) (Input, string) {
	var id int
	var password string
	var imageUrl string
	var in Input
	row, err := db.Query("select userid,password,image from users where email=$1", o.Email)
	defer row.Close()

	check(err)
	if row.Next() {
		err = row.Scan(&id, &password, &imageUrl)
		check(err)
		fmt.Println(password, "2")

		if password == o.Password {
			//放回登陆成功
			in.ID = strconv.Itoa(id)
			in.Token = GetToken(in.ID)
			in.IfLogin = true
			in.Tip = "登陆成功"
		} else {
			//返回登陆密码错误
			in.IfLogin = false
			in.Tip = "密码错误"
		}
	} else {
		//返回账号不存在
		in.IfLogin = false
		in.Tip = "Email不存在"
	}
	return in, imageUrl
} //登录-查错-

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

//检查令牌
func CheckToken(token string, k string) bool {
	key := []byte(k)
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return false
	}
	return true
}

//得到令牌
func GetToken(k string) string {
	key := []byte(k)
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + (60 * 60 * 2)), //秒为单位
		Issuer:    "login",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return ss
}

//注册检验是否存在用户
func Register(re *HeadUser) {
	dbuser := HeadUser{} //数据库数据载体
	rows, err := db.Query("SELECT Email from users")
	if err != nil {
		fmt.Println("*1*\n")
		check(err)
	}
	for rows.Next() {
		err = rows.Scan(&dbuser.Email)
		if dbuser.Email == re.Email {
			re.Email = "0"
			fmt.Println("用户已经存在")
		}
	}
	rows.Close()

	// return re
}

//注册数据插入数据库
func RegisterInsert(re *HeadUser) {
	stmt, err := db.Prepare("INSERT INTO users( Email,Password,image,registrationdate) values($1,$2,$3,$4)")
	if err != nil {
		//err
		fmt.Println("*3*\n:")
		check(err)
	}

	timeNow := time.Now()
	Imageurl := "./78dea7ef1543c888ee43ded56d49ed2.png"
	_, err = stmt.Exec(re.Email, re.Password, Imageurl, timeNow)
	if err != nil {
		//err
		fmt.Println("*4*\n")
	} else {
		fmt.Println("insert successed!")
	}
}

//修改密码
func ChangePassword(re *HeadUser) {
	rows, err := db.Query("SELECT Email from users where email=$1", re.Email)
	if err != nil {
		fmt.Println("*1*\n")
		check(err)
	}
	if rows.Next() { //若存在此账号则修改数据库数据
		stmt, err := db.Prepare("update users set password=$1")
		check(err)
		_, err = stmt.Exec(re.Password)
		check(err)
	} else {
		if !strings.Contains(re.Email, "@") {
			re.Email = "1" //非邮件格式（不含@）
		} else {
			re.Email = "0" //用户不存在
		}
	}
	rows.Close()

}
