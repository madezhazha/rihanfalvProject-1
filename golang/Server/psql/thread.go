package psql

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
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
	rows, err := db.Query("SELECT topicid,posterid,topictitle,topiccontent,creationtime,numberofreplies,finalreplytime,collectionvolume,visitvolume,japanorkorea FROM topics ORDER BY topicid desc")
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
	err = db.QueryRow("SELECT userid, username, email, image FROM users WHERE userid = $1", thread.Userid).
		Scan(&user.Userid, &user.Username, &user.Email, &user.Image)

	// 如果不包含“assets”，转成base64.如果包含，直接传该地址
	if !strings.Contains(user.Image, "assets") {
		user.Image = ImgToBase64(user.Image) // 通过路径读取图片，并转成base64传给前端
	}
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

// AddCollectNum 主贴的收藏数加一
func AddCollectNum(topicID int) (err error) {
	stmt, err := db.Prepare("update topics set collectionvolume=collectionvolume+1 where topicid=$1")
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

// CutCollectNum 主贴的收藏数减一
func CutCollectNum(topicID int) (err error) {
	stmt, err := db.Prepare("update topics set collectionvolume=collectionvolume-1 where topicid=$1")
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

// AddVisitNum 浏览数加一
func AddVisitNum(topicID int) (err error) {
	stmt, err := db.Prepare("update topics set visitvolume=visitvolume+1 where topicid=$1")
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

// UpdFinReplyTime 更新最后主贴回复时间
func UpdFinReplyTime(topicID int) (err error) {
	stmt, err := db.Prepare("update topics set finalreplytime=$1 where topicid=$2")
	if err != nil {
		fmt.Println("Prepare:", err)
		return err
	}

	_, err = stmt.Exec(time.Now(), topicID)
	if err != nil {
		fmt.Println("Exec:", err)
		return err
	}
	return err
}

// RsByCondition 根据查询条件查询出结果
func RsByCondition(condition []string) (usersAndThreads []map[string]interface{}, err error) {
	for i, value := range condition {
		rows, err := db.Query("SELECT topicid,posterid,topictitle,topiccontent,creationtime FROM topics where label like $1 ORDER BY topicid", "%"+value+"%")
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			conv := Thread{}
			user := MyUser{}
			userAndThread := make(map[string]interface{})
			if err = rows.Scan(&conv.ID, &conv.Userid, &conv.Topictitle, &conv.Topiccontent, &conv.Creationtime); err != nil {
				return nil, err
			}
			user, err = conv.User()
			if i == 0 {
				userAndThread["thread"] = conv
				userAndThread["user"] = user
				usersAndThreads = append(usersAndThreads, userAndThread)
			} else {
				// 判断usersAndThreads里面是否已经包含了查询结果的标志
				flag := true
				for _, val := range usersAndThreads {
					// 类型断言，因为val[string]本身是interface{}类型，没有属性.
					// 所以即使确定了key，即现在为val["thread"]在编译时期的类型也会被识别为interface{}.
					// 解决办法就是使用类型断言，把现在类型便转换成Thread(本身语法不允许使用强制类型转换)，而Thread才有属性
					v, ok := val["thread"].(Thread)
					if ok {
						if v.ID == conv.ID {
							flag = false
						}
					} else {
						return nil, err
					}
				}
				if flag == true {
					userAndThread["thread"] = conv
					userAndThread["user"] = user
					usersAndThreads = append(usersAndThreads, userAndThread)
				}
			}
		}
	}
	return
}

// ImgToBase64 根据图片的地址得到图片的二进制，再把图片转成字符串（抄袭组长的函数）
func ImgToBase64(image string) string {
	img, _ := ioutil.ReadFile(image)              //直接读取文件内容，内容是[]byte类型
	str := base64.StdEncoding.EncodeToString(img) //将[]byte转化成string
	return str
}
