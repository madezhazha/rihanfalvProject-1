package psql

import (
	"strings"
	"time"
)

// Post 回帖的结构体
type Post struct {
	Replieid     int
	Userid       int
	Topicid      int
	Repiycontent string
	Floor        int
	Replytime    JSONTime
}

// Collection 收藏的结构体
type Collection struct {
	Collectionid        int
	Userid              int
	Collectioncontentid int
	Collectiontime      JSONTime
	Collectiontype      string
}

// GetPost 根据主贴id获取主贴所有的回帖(包过用户)
func GetPost(topicid int) (usersAndPosts []map[string]interface{}, err error) {
	rows, err := db.Query("SELECT * FROM replies WHERE topicid=$1 ORDER BY replieid", topicid)
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Post{}
		user := MyUser{}
		userAndPost := make(map[string]interface{})
		if err = rows.Scan(
			&conv.Replieid, &conv.Userid, &conv.Topicid, &conv.Repiycontent,
			&conv.Floor, &conv.Replytime); err != nil {
			return
		}
		user, err = conv.User()
		if err != nil {
			return
		}

		userAndPost["post"] = conv
		userAndPost["user"] = user
		usersAndPosts = append(usersAndPosts, userAndPost)
	}
	rows.Close()
	return
}

// User 根据回帖的userid找到用户
func (post *Post) User() (user MyUser, err error) {
	err = db.QueryRow("SELECT userid, username, email, image FROM users WHERE userid = $1", post.Userid).
		Scan(&user.Userid, &user.Username, &user.Email, &user.Image)
	if err != nil {
		return
	}

	// 如果不包含“assets”，转成base64.如果包含，直接传该地址
	if !strings.Contains(user.Image, "assets") {
		user.Image = ImgToBase64(user.Image) // 通过路径读取图片，并转成base64传给前端
	}
	return
}

// CreatePost 新建一个新的回帖
func CreatePost(userID int, topicID int, text string, floor int) error {
	statement := "insert into replies (userid,topicid,replycontent,floor,replytime) values ($1,$2,$3,$4,$5) "
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, topicID, text, floor, time.Now())
	if err != nil {
		return err
	}

	return err
}

// IsCollected 根据用户id和主贴id判断收藏表里面指定用户是否收藏某个主贴
func IsCollected(userID int, topicID int) (collection Collection, err error) {
	rows, err := db.Query("SELECT * FROM collection WHERE userid = $1 AND collectiontype=$2 AND collectioncontentid=$3",
		userID, "topics", topicID)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&collection.Collectionid, &collection.Userid, &collection.Collectiontype, &collection.Collectioncontentid,
			&collection.Collectiontime); err != nil {
			return
		}
	}
	return
}

// Collect 指定用户收藏指定主贴并返回收藏表的id
func Collect(userID int, topicID int) (collectionid int, err error) {
	statement := "insert into collection (userid, collectioncontentid,collectiontime, collectiontype) values ($1,$2,$3,$4) returning collectionid"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(userID, topicID, time.Now(), "topics").Scan(&collectionid)
	return
}

// Cancel 指定用户取消收藏指定主贴
func Cancel(userID int, topicID int) (collectionid int, err error) {
	statement := "DELETE FROM collection WHERE userid=$1 and collectioncontentid = $2 and collectiontype=$3"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, topicID, "topics")
	if err != nil {
		return
	}
	collectionid = 0
	return
}
