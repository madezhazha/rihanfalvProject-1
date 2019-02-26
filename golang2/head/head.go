package head

import (
	 "database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func oninit() {
	var err error
	db, err = sql.Open("postgres", "user=User1 password=123456 dbname=JAndKLaw sslmode=disable")
	//defer db.Close()
	check(err)
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	dbuser := User{} //数据库数据载体
	re := User{}     //前端数据载体
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//err
		check(err)
	}
	err = json.Unmarshal([]byte(body), &re)
	if err != nil {
		//err
		check(err)
	}
	fmt.Println("\ngetdata:", re)

	//查询数据是否已经注册在数据库
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

	if re.Email != "0" {
		//用户不存在，数据插入数据库
		if !strings.Contains(re.Email, "@") {
			re.Email = "1"
		} else {
			if strings.Contains(re.Password, " ") {
				re.Email = "2"
			} else {
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
		}

	}

	b, err := json.Marshal(re)
	if err != nil {
		fmt.Println("*5*\n")
		fmt.Println("enconding faild")
	} else {
		fmt.Println("encoding successful")
		fmt.Println(string(b))
	}
	w.Write(b)

} //注册功能

func Email(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	re := User{}

	//获取邮箱
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//err
		check(err)
	}
	err = json.Unmarshal([]byte(body), &re)
	if err != nil {
		//err
		check(err)
	}
	fmt.Println("\ngetdata:", re)

	//若是邮箱格式就进行发送
	var RandNumberString string
	if strings.Contains(re.Email, "@") { //是邮箱格式
		Bool := true
		rand.Seed(time.Now().UnixNano())
		for Bool {
			randNumberInt := rand.Intn(1000000)
			if randNumberInt >= 100000 {
				Bool = false                                   //用于判断循环结束
				RandNumberString = strconv.Itoa(randNumberInt) //整型转化为字符串型

				auth := smtp.PlainAuth("", "1062111902@qq.com", "ehzwrwocsocubfia", "smtp.qq.com")
				to := []string{re.Email}
				nickname := "验证码"
				user := "1062111902@qq.com"
				subject := "验证码"
				content_type := "Content-Type: text/plain; charset=UTF-8"
				body := "【日韩法律web应用】欢迎使用本应用，您的验证码是：" + RandNumberString + "。如非本人操作，请注意账号安全。"
				msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
					"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
				err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
				if err != nil {
					fmt.Printf("send mail error: %v", err)
					RandNumberString = "1"
				} else {
					fmt.Println("send eamil successed! the Verification is :", RandNumberString)
				}
			}
		}
	} else {
		RandNumberString = "0" //若不是邮箱（含@）的格式
	}

	//传值回前端
	b, err := json.Marshal(RandNumberString)
	if err != nil {
		fmt.Println("enconding faild")
	} else {
		fmt.Println("encoding successful ")
		fmt.Println(string(b))
	}
	w.Write(b)
} //发送验证码

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	re := User{} //前端数据载体
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//err
		check(err)
	}
	err = json.Unmarshal([]byte(body), &re)
	if err != nil {
		//err
		check(err)
	}
	fmt.Println("\ngetdata:", re)

	//修改数据库
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

	b, err := json.Marshal(re)
	if err != nil {
		fmt.Println("*5*\n")
		fmt.Println("enconding faild")
	} else {
		fmt.Println("encoding successful")
		fmt.Println(string(b))
	}
	w.Write(b)

} //修改密码

func CPsendVerification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	re := User{}
	//获取邮箱
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//err
		check(err)
	}
	err = json.Unmarshal([]byte(body), &re)
	if err != nil {
		//err
		check(err)
	}
	fmt.Println("\ngetdata:", re)

	//若是邮箱格式就进行发送
	var RandNumberString string
	if strings.Contains(re.Email, "@") { //是邮箱格式
		Bool := true
		rand.Seed(time.Now().UnixNano())
		for Bool {
			randNumberInt := rand.Intn(1000000)
			if randNumberInt >= 100000 {
				Bool = false                                   //用于判断循环结束
				RandNumberString = strconv.Itoa(randNumberInt) //整型转化为字符串型

				auth := smtp.PlainAuth("", "1062111902@qq.com", "ehzwrwocsocubfia", "smtp.qq.com")
				to := []string{re.Email}
				nickname := "验证码"
				user := "1062111902@qq.com"
				subject := "验证码"
				content_type := "Content-Type: text/plain; charset=UTF-8"
				body := "【日韩法律web应用】欢迎使用本应用，您的验证码是：" + RandNumberString + "。如非本人操作，请注意账号安全。"
				msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
					"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
				err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
				if err != nil {
					fmt.Printf("send mail error: %v", err)
					RandNumberString = "1"
				} else {
					fmt.Println("send eamil successed! the Verification is :", RandNumberString)
				}
			}
		}
	} else {
		RandNumberString = "0" //若不是邮箱（含@）的格式
	}

	//传值回前端
	b, err := json.Marshal(RandNumberString)
	if err != nil {
		fmt.Println("enconding faild")
	} else {
		fmt.Println("encoding successful ")
		fmt.Println(string(b))
	}
	w.Write(b)
} //发送修改密码的验证码


func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

/*数据库建表语句
create table Users(
	"userid" serial primary key,
	"userName" char(20),
	"password" char(20),
	"email" char(30),
	"image" char(100),
	"integal"  integer,
	"registrationdate" char(100)

	)
*/


// func main() {
// 	connectdb()
// 	server := http.Server{
// 		Addr: "127.0.0.1:8090",
// 	}
// 	http.HandleFunc("/", log(hello))
// 	http.HandleFunc("/register", register)
// 	http.HandleFunc("/sendVerification", email)
// 	http.HandleFunc("/CPsendVerification", CPsendVerification)
// 	http.HandleFunc("/changePassword", changePassword)
// 	server.ListenAndServe()
// }
