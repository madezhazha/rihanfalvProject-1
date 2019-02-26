package chat


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "sql1234567"
	dbname   = "translateinfo"
)

//用户信息
type Users struct {
	Userid   int    `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Integral int    `json:"integral"`
}

//主贴
type Topics struct {
	Topicid          int    `json:"topicid"`
	Userid           int    `json:"userid"`
	Topictitle       string `json:"topictitle"`
	Topiccontent     string `json:"topiccontent"`
	Numberofreplies  int    `json:"numberofreplies"`
	Collectionvolume int    `json:"collectionvolume"`
	Visitvolume      int    `json:"visitvolume"`
	Japanorkorea     int    `json:"japanorkorea"`
}

//回帖
type Replies struct {
	Replieid     int    `json:"replieid"`
	Userid       int    `json:"userid"`
	Topicid      int    `json:"topicid"`
	Replycontent string `json:"replycontent"`
	Floor        int    `json:"floor"`
}

// 通过用户id获取信息
type Getid struct {
	Userid int `json:"userid"`
}