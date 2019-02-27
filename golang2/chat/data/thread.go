package data

import (
	"ChatServe/utils"
	"database/sql"
	"fmt"
	"log"

	// 驱动包
	_ "github.com/lib/pq"
)

// Db 全局变量
var Db *sql.DB

// Thread 帖子的结构体
type Thread struct {
	ID               int
	Userid           int
	Topictitle       string
	Topiccontent     string
	Creationtime     utils.JSONTime
	Numberofreplies  int
	Finalerplytime   utils.JSONTime
	Collectionvolume int
	Visitvolume      int
	Japanorkorea     int
	Tag              string
}

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres password=1234 dbname=law sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}

// GetThread 从数据库获取所有的帖子
func GetThread() (usersAndThreads []map[string]interface{}, err error) {
	rows, err := Db.Query("SELECT topicid,userid,topictitle,topiccontent,creationtime,numberofreplies,finalreplytime,collectionvolume,visitvolume,japanorkorea FROM topics ORDER BY topicid")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		user := User{}
		userAndThread := make(map[string]interface{})
		if err = rows.Scan(
			&conv.ID, &conv.Userid, &conv.Topictitle, &conv.Topiccontent, &conv.Creationtime, &conv.Numberofreplies,
			&conv.Finalerplytime, &conv.Collectionvolume, &conv.Visitvolume, &conv.Japanorkorea); err != nil {
			return
		}
		user, err = conv.User()
		if err != nil {
			fmt.Println("err")
			return
		}
		userAndThread["thread"] = conv
		userAndThread["user"] = user
		usersAndThreads = append(usersAndThreads, userAndThread)
	}
	rows.Close()
	return
}

// User 根据主帖的userid找到指定的用户
func (thread *Thread) User() (user User, err error) {
	err = Db.QueryRow("SELECT userid, username, email, image,integral,registrationdate FROM users WHERE userid = $1", thread.Userid).
		Scan(&user.Userid, &user.Username, &user.Email, &user.Image, &user.Integral, &user.Registrationdate)
	return
}

// ThreadByTopicID 根据topicid找到对应的唯一主贴(包过用户)
func ThreadByTopicID(topicID int) (userAndThread map[string]interface{}, err error) {
	thread := Thread{}
	userAndThread = make(map[string]interface{})
	err = Db.QueryRow("SELECT topicid,userid,topictitle,topiccontent,creationtime FROM topics WHERE topicid = $1", topicID).
		Scan(&thread.ID, &thread.Userid, &thread.Topictitle, &thread.Topiccontent, &thread.Creationtime)
	userAndThread["thread"] = thread
	userAndThread["user"], err = thread.User()
	if err != nil {
		return nil, err
	}
	return
}

// AddRepNum 用户评论数加一
func AddRepNum(topicID int) error {
	stmt, err := Db.Prepare("update topics set numberofreplies=numberofreplies+1 where topicid=$1")
	if err != nil {
		fmt.Println("Prepare:", err)
		return err
	}

	_, err = stmt.Exec(topicID)
	if err != nil {
		fmt.Println("Exec:", err)
		return err
	}
	return err
}

// SetText 把文本设置进数据库
func SetText(text string) error {
	sql := "update replies set replycontent = $1 where replieid = 4"
	// sql := "insert into test (thread) values ($1) returning id,thread"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(text)
	if err != nil {
		return err
	}
	return nil
}
