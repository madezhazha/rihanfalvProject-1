package psql

import (
	"fmt"
	"time"

	// 驱动包
	_ "github.com/lib/pq"
)

// JSONTime json的格式化时间
type JSONTime time.Time

// MyUser 用户的结构体
type MyUser struct {
	Userid           int
	Username         string
	Password         string
	Email            string
	Image            string
	Integral         int
	Registrationdate JSONTime
}

// Thread 帖子的结构体
type Thread struct {
	ID               int
	Userid           int
	Topictitle       string
	Topiccontent     string
	Creationtime     JSONTime
	Numberofreplies  int
	Finalerplytime   JSONTime
	Collectionvolume int
	Visitvolume      int
	Japanorkorea     int
	label            string
}

// GetThread 从数据库获取所有的帖子
func GetThread() (usersAndThreads []map[string]interface{}, err error) {
	rows, err := db.Query("SELECT topicid,posterid,topictitle,topiccontent,creationtime,numberofreplies,finalreplytime,collectionvolume,visitvolume,japanorkorea FROM topics ORDER BY topicid")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		user := MyUser{}
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
func (thread *Thread) User() (user MyUser, err error) {
	err = db.QueryRow("SELECT userid, username, email, image,integral,registrationdate FROM users WHERE userid = $1", thread.Userid).
		Scan(&user.Userid, &user.Username, &user.Email, &user.Image, &user.Integral, &user.Registrationdate)
	return
}

// ThreadByTopicID 根据topicid找到对应的唯一主贴(包过主贴的用户信息)
func ThreadByTopicID(topicID int) (userAndThread map[string]interface{}, err error) {
	thread := Thread{}
	userAndThread = make(map[string]interface{})
	err = db.QueryRow("SELECT topicid,posterid,topictitle,topiccontent,creationtime FROM topics WHERE topicid = $1", topicID).
		Scan(&thread.ID, &thread.Userid, &thread.Topictitle, &thread.Topiccontent, &thread.Creationtime)
	if err != nil {
		return
	}
	userAndThread["thread"] = thread
	userAndThread["user"], err = thread.User()
	if err != nil {
		return nil, err
	}
	return
}

// AddRepNum 用户评论数加一
func AddRepNum(topicID int) error {
	stmt, err := db.Prepare("update topics set numberofreplies=numberofreplies+1 where topicid=$1")
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

// RsByCondition 根据查询条件查询出结果
func RsByCondition(condition []string) (thread []Thread, err error) {
	for i, value := range condition {
		rows, err := db.Query("SELECT topicid,posterid,topictitle,topiccontent,creationtime FROM topics where label like $1 ORDER BY topicid", "%"+value+"%")
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			conv := Thread{}
			if err = rows.Scan(&conv.ID, &conv.Userid, &conv.Topictitle, &conv.Topiccontent, &conv.Creationtime); err != nil {
				return nil, err
			}
			if i == 0 {
				thread = append(thread, conv)
			} else {
				flag := true
				for _, value := range thread {
					if value.ID == conv.ID {
						flag = false
					}
				}
				if flag == true {
					thread = append(thread, conv)
				}
			}
		}
	}
	return
}
