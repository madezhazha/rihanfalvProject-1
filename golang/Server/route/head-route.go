package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"../psql"

	_ "github.com/lib/pq"
)

func Log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")             //返回数据格式是json
		h(w, r)
	}
} //跨域是函数头

func Hello(w http.ResponseWriter, r *http.Request) {
	var in psql.Input
	var imageUrl string

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	headcheck(err)
	fmt.Println(string(body), "******")
	var u psql.Output
	err = json.Unmarshal(body, &u)
	in, imageUrl = psql.CheckEmailPassword(u)

	if in.IfLogin {
		img, err := ioutil.ReadFile(imageUrl)
		if err != nil {
			fmt.Println("read file error")
		}
		in.Image = img
	}
	data, _ := json.Marshal(in)
	w.Write(data)
} //登录。。。

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	// dbuser := psql.User{} //数据库数据载体
	re := psql.HeadUser{} //前端数据载体
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//err
		headcheck(err)
	}
	err = json.Unmarshal([]byte(body), &re)
	if err != nil {
		//err
		headcheck(err)
	}
	fmt.Println("\ngetdata:", re)

	//查询数据是否已经注册在数据库
	psql.Register(&re)
	/*	rows, err := psql.Db.Query("SELECT Email from users")
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
		rows.Close()*/

	if re.Email != "0" {
		//用户不存在，数据插入数据库
		if !strings.Contains(re.Email, "@") {
			re.Email = "1"
		} else {
			if strings.Contains(re.Password, " ") {
				re.Email = "2"
			} else {
				psql.RegisterInsert(&re)
				/*stmt, err := psql.Db.Prepare("INSERT INTO users( Email,Password,image,registrationdate) values($1,$2,$3,$4)")
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
				}*/
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

	re := psql.HeadUser{}

	//获取邮箱
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//err
		headcheck(err)
	}
	err = json.Unmarshal([]byte(body), &re)
	if err != nil {
		//err
		headcheck(err)
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

	re := psql.HeadUser{} //前端数据载体
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//err
		headcheck(err)
	}
	err = json.Unmarshal([]byte(body), &re)
	if err != nil {
		//err
		headcheck(err)
	}
	fmt.Println("\ngetdata:", re)

	//修改数据库
	psql.ChangePassword(&re)
	/*	rows, err := psql.Db.Query("SELECT Email from users where email=$1", re.Email)
		if err != nil {
			fmt.Println("*1*\n")
			check(err)
		}
		if rows.Next() { //若存在此账号则修改数据库数据
			stmt, err := psql.Db.Prepare("update users set password=$1")
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
		rows.Close()*/

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

	re := psql.HeadUser{}
	//获取邮箱
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//err
		headcheck(err)
	}
	err = json.Unmarshal([]byte(body), &re)
	if err != nil {
		//err
		headcheck(err)
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

//检查错误
func headcheck(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
