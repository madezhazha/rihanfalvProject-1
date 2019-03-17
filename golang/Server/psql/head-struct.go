package psql

type HeadUser struct {
	Email    string
	Password string
} //注册的信息

type Input struct {
	ID       string
	IfLogin  bool
	Tip      string
	Image    []byte
	Token    string
	ImageUrl string
} //输入数据结构

type Output struct {
	Email    string
	Password string
} //输出数据结构
