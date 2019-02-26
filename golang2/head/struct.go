package head

type User struct {
	Email    string
	Password string
} //注册的信息

type input struct {
	IfLogin bool
	Tip     string
	Image   []byte
} //输入数据结构

type output struct {
	Email    string
	Password string
} //输出数据结构
