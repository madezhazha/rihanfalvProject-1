package psql

import (
	"encoding/base64"
	"strings"
	"database/sql"

	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	//"net/http"
	"time"

	_ "github.com/lib/pq"
)

//用户信息
type Users struct {
	Userid   int    `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Image	string `json:"image"`
	Integral int    `json:"integral"`
}

//主贴
type Topics struct {
	Topicid          int       `json:"topicid"`
	Userid           int       `json:"userid"`
	Topictitle       string    `json:"topictitle"`
	Topiccontent     string    `json:"topiccontent"`
	Creationtime     time.Time `json:"creationtime"`
	Numberofreplies  int       `json:"numberofreplies"`
	Finalreplytime   time.Time `json:"finalreplytime"`
	Collectionvolume int       `json:"collectionvolume"`
	Visitvolume      int       `json:"visitvolume"`
	Japanorkorea     int       `json:"japanorkorea"`
	Topiclabel       string    `json:"topiclabel"`
}

//回帖
type Replies struct {
	Replieid     int       `json:"replieid"`
	Userid       int       `json:"userid"`
	Topicid      int       `json:"topicid"`
	Replycontent string    `json:"replycontent"`
	Floor        int       `json:"floor"`
	Replytime    time.Time `json:"replytime"`
}

// 通过用户id获取信息
type Getid struct {
	Userid int `json:"userid"`
}


//获取个人信息
func GetUserInfo(postbody Users) Users {
	getInfo := " select username,password,email,image,integral from users where userid=$1"
	err := db.QueryRow(getInfo, &postbody.Userid).Scan(&postbody.Username, &postbody.Password, &postbody.Email, &postbody.Image, &postbody.Integral)
	if err != nil {
		if err == sql.ErrNoRows {
		} else {
			log.Fatal(err)
		}
	}

	// 如果不包含“assets”，转成base64.如果包含，直接传该地址
	if !strings.Contains(postbody.Image, "assets") {
		postbody.Image = Img_ToBase64(postbody.Image) // 通过路径读取图片，并转成base64传给前端
	}

	return postbody
}

//获取主贴
func GetUserTopic(postbody int) []Topics {
	var usertopics []Topics
	getTopic  := " select topicid,posterid,topictitle,topiccontent,creationtime - interval '8 Hours',numberofreplies,finalreplytime - interval '8 Hours',collectionvolume,visitvolume,japanorkorea,label from topics where posterid=$1 "
	rows, err := db.Query(getTopic, postbody)
	if err != nil {
		fmt.Println(err)
		//return
	}
	defer rows.Close()
	for rows.Next() {
		var temp Topics
		err := rows.Scan(&temp.Topicid, &temp.Userid, &temp.Topictitle, &temp.Topiccontent, &temp.Creationtime, &temp.Numberofreplies, &temp.Finalreplytime, &temp.Collectionvolume, &temp.Visitvolume, &temp.Japanorkorea, &temp.Topiclabel)

		if err != nil {
			fmt.Println(err)
			//return
		}

		usertopics = append(usertopics, temp)
		//fmt.Println("Test replytime :", temp.Finalreplytime)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	return usertopics
}

//获取回帖
func GetUserReply(postbody Getid) []Replies {
	var userreplies []Replies
	getReply := " select replieid,userid,topicid,replycontent,floor,replytime - interval '8 Hours' from replies where userid=$1 "
	rows, err1 := db.Query(getReply, postbody.Userid)
	if err1 != nil {
		fmt.Println(err1)
		//return
	}
	defer rows.Close()
	for rows.Next() {
		var temp Replies
		err := rows.Scan(&temp.Replieid, &temp.Userid, &temp.Topicid, &temp.Replycontent, &temp.Floor, &temp.Replytime)

		if err != nil {
			fmt.Println(err)
			//return
		}

		userreplies = append(userreplies, temp)
	}
	err := rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	return userreplies
}

//添加主贴
func InsertTopic(postbody Topics) {
	insertTopic := "insert into topics(posterid,topictitle,topiccontent,japanorkorea,label) values($1,$2,$3,$4,$5)"
	stmt, err := db.Prepare(insertTopic)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(postbody.Userid, postbody.Topictitle, postbody.Topiccontent, postbody.Japanorkorea, postbody.Topiclabel)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert into user_tbl success")
	}
}

//阅读量加一
func AddTopicVisNum(topicId int) error {
	addVisNum := "update topics set visitvolume=visitvolume+1 where topicid=$1"
	stmt, err := db.Prepare(addVisNum)
	if err != nil {
		fmt.Println("Prepare:", err)
		return err
	}
	_, err = stmt.Exec(topicId)
	if err != nil {
		fmt.Println("Exec:", err)
		return err
	} else {
		fmt.Println("udpate topics-visitvolume success")
	}
	return nil
}

// Img_ToBase64 根据图片的地址得到图片的二进制，再把图片转成字符串（抄袭组长以及队友的函数）
func Img_ToBase64(image string) string {
	img, _ := ioutil.ReadFile(image)              //直接读取文件内容，内容是[]byte类型
	str := base64.StdEncoding.EncodeToString(img) //将[]byte转化成string
	return str
}