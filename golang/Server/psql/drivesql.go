package psql

import(
	"fmt"
	"strings"
	"strconv"
)



//创建一个结构体用来接收所有的数据
type Information struct{
	ID string
	trialtime string
	causetype string 
	concretecasetype string
	judgetype string
	trialgrade string
	title string
	event string
	Type string
}


type payment struct{
	payId string
	userId int
	pointID int
}



type concretelaw1 struct{
	ID string 
	viewpoint string
	casetitle string
	header string
	Type string
	integral int
}

type collectioncase struct{
	ID string
	collectiontype string
	collectiontitle string
	collectiontitleid string
	userId string
}

type collection struct{
	userid int   //用户id
	collectionid int   //收藏的id
	collectiontype string //收藏该文章的语言类型
	collectioncontentid int //收藏该文章的id，就相当于titleid
	collectiontime string //收藏时间
}




//获取所有的数据
func Getalldata(languageType string,NumberCasethingString string)interface{}{

	all_data := make(map[string]map[string]string)

	//数据库的查询
	rows,err := db.Query("select * from casething where type=$1 limit 5 offset $2",languageType,NumberCasethingString)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event,&info.Type)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误"
		}

		//将数据放在了一块
		all_data[info.ID] = make(map[string]string)
		all_data[info.ID]["id"]=info.ID
		all_data[info.ID]["trialtime"] = info.trialtime
		all_data[info.ID]["causeofaction"] = info.causetype
		all_data[info.ID]["concretecasetype"] = info.concretecasetype
		all_data[info.ID]["legalprinciple"] = info.judgetype
		all_data[info.ID]["trialgrade"] = info.trialgrade
		all_data[info.ID]["casetitle"] = info.title
		all_data[info.ID]["casecontent"] = info.event
		all_data[info.ID]["type"] = info.Type

	}

	rows.Close()

	return all_data
}

func Getfirstfloor(data string,languageType string,NumberCasethingString string)interface{}{
	all_data := make(map[string]map[string]string)

	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from casething where causeofaction=$1 and type=$2 limit 5 offset $3",data,languageType,NumberCasethingString)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event,&info.Type)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误"
		}

		//将数据放在了一块
		all_data[info.ID] = make(map[string]string)
		all_data[info.ID]["id"]=info.ID
		all_data[info.ID]["trialtime"] = info.trialtime
		all_data[info.ID]["causeofaction"] = info.causetype
		all_data[info.ID]["concretecasetype"] = info.concretecasetype
		all_data[info.ID]["legalprinciple"] = info.judgetype
		all_data[info.ID]["trialgrade"] = info.trialgrade
		all_data[info.ID]["casetitle"] = info.title
		all_data[info.ID]["casecontent"] = info.event
		all_data[info.ID]["type"] = info.Type
	}

	rows.Close()

	return all_data
}


func Getreason(data string,languageType string,NumberCasethingString string)interface{}{
	all_data := make(map[string]map[string]string)

	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from casething where concretecasetype=$1 and type=$2 limit 5 offset $3",data,languageType,NumberCasethingString)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event,&info.Type)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误"
		}

		//将数据放在了一块
		all_data[info.ID] = make(map[string]string)
		all_data[info.ID]["id"]=info.ID
		all_data[info.ID]["trialtime"] = info.trialtime
		all_data[info.ID]["causeofaction"] = info.causetype
		all_data[info.ID]["concretecasetype"] = info.concretecasetype
		all_data[info.ID]["legalprinciple"] = info.judgetype
		all_data[info.ID]["trialgrade"] = info.trialgrade
		all_data[info.ID]["casetitle"] = info.title
		all_data[info.ID]["casecontent"] = info.event
		all_data[info.ID]["type"] = info.Type
	}

	rows.Close()

	return all_data
}


func Getlevel(data string,languageType string,NumberCasethingString string)interface{}{
	all_data := make(map[string]map[string]string)

	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from casething where trialgrade=$1 and type=$2 limit 5 offset $3",data,languageType,NumberCasethingString)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event,&info.Type)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误"
		}

		//将数据放在了一块
		all_data[info.ID] = make(map[string]string)
		all_data[info.ID]["id"]=info.ID
		all_data[info.ID]["trialtime"] = info.trialtime
		all_data[info.ID]["causeofaction"] = info.causetype
		all_data[info.ID]["concretecasetype"] = info.concretecasetype
		all_data[info.ID]["legalprinciple"] = info.judgetype
		all_data[info.ID]["trialgrade"] = info.trialgrade
		all_data[info.ID]["casetitle"] = info.title
		all_data[info.ID]["casecontent"] = info.event
		all_data[info.ID]["type"] = info.Type
	}

	rows.Close()

	return all_data
} 

func Getsecondfloor(data string,languageType string,NumberCasethingString string)interface{}{
	all_data := make(map[string]map[string]string)


	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from casething where legalprinciple=$1 and type=$2 limit 5 offset $3",data,languageType,NumberCasethingString)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event,&info.Type)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误"
		}

		//将数据放在了一块
		all_data[info.ID] = make(map[string]string)
		all_data[info.ID]["id"]=info.ID
		all_data[info.ID]["trialtime"] = info.trialtime
		all_data[info.ID]["causeofaction"] = info.causetype
		all_data[info.ID]["concretecasetype"] = info.concretecasetype
		all_data[info.ID]["legalprinciple"] = info.judgetype
		all_data[info.ID]["trialgrade"] = info.trialgrade
		all_data[info.ID]["casetitle"] = info.title
		all_data[info.ID]["casecontent"] = info.event
		all_data[info.ID]["type"] = info.Type
	}

	rows.Close()

	return all_data
}

func Gettime(data string,languageType string,NumberCasethingString string)interface{}{
	all_data := make(map[string]map[string]string)

	a:=0

	rows,err := db.Query("select * from casething where type=$1 limit 5 offset $2",languageType,NumberCasethingString)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event,&info.Type)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误"
		}

		//这里要判断才能将数据放进入
		year := strings.Split(info.trialtime,"/")

		for i:=0;i<3;i++{
			if year[i]==data{
				a=1;
			}
		}
		if a==1{
			all_data[info.ID] = make(map[string]string)
			all_data[info.ID]["id"]=info.ID
			all_data[info.ID]["trialtime"] = info.trialtime
			all_data[info.ID]["causeofaction"] = info.causetype
			all_data[info.ID]["concretecasetype"] = info.concretecasetype
			all_data[info.ID]["legalprinciple"] = info.judgetype
			all_data[info.ID]["trialgrade"] = info.trialgrade
			all_data[info.ID]["casetitle"] = info.title
			all_data[info.ID]["casecontent"] = info.event
			all_data[info.ID]["type"] = info.Type
			a=0;
		}
	}

	rows.Close()

	return all_data

}


//法官观点,没有用户的获取
func Gettext_nouserid(data string)interface{}{
	var con concretelaw1

	//创建一个单个集合的承接数据
	all:=make(map[string]string)
	rows,err:=db.Query("select * from point where casetitle=$1",data) 

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}
	for rows.Next(){
		err:=rows.Scan(&con.ID,&con.viewpoint,&con.casetitle,&con.header,&con.Type,&con.integral)
		if err!=nil{
			fmt.Println(err)
			return "系统出现错误"
		}
		all["viewpoint"]=con.viewpoint
		all["casetitle"]=con.casetitle
		all["header"]=con.header
		all["type"]=con.Type
		all["ID"]=con.ID
		all["integral"]=strconv.Itoa(con.integral)  //将它转化为字符串
	}

	defer rows.Close()

	return all
}

//获取法官观点，有用户
func Gettext_userid(data string,userId string)interface{}{
	var con concretelaw1

	//创建一个单个集合的承接数据
	all:=make(map[string]string)
	rows,err:=db.Query("select * from point where casetitle=$1",data) 

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}
	for rows.Next(){
		err:=rows.Scan(&con.ID,&con.viewpoint,&con.casetitle,&con.header,&con.Type,&con.integral)
		if err!=nil{
			fmt.Println(err)
			return "系统出现错误"
		}
		all["viewpoint"]=con.viewpoint
		all["casetitle"]=con.casetitle
		all["header"]=con.header
		all["type"]=con.Type
		all["ID"]=con.ID
		all["integral"]=strconv.Itoa(con.integral)  //将它转化为字符串
	}

	//默认title是唯一的,查找是否已经付钱了
	searchResult:=searchpay(con.ID,userId)
	all["searchResult"] = searchResult
	//放回用户的积分数
	integral:=GEtintegral(userId)
	all["allintergral"]=integral
	//获取用户名
	username:=Getusername(userId)
	all["username"]=username
	defer rows.Close()
	return all
}

//利用该文章的id和用户的id进行查找，是否是付费的顾客
func searchpay(titleId string,userId string)string{
	var pay payment
	//将字符串转化为数字
	userID,err2 := strconv.Atoi(userId)
	if(err2!=nil){
		fmt.Println("titleid或userid的转化的结果出现错误",err2)
		return "系统出现错误"
	}
	rows,err:=db.Query("select * from payment where pointid=$1",titleId)
	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}
	for rows.Next(){
		err:=rows.Scan(&pay.payId,&pay.userId,&pay.pointID)
		if err!=nil{
			fmt.Println(err)
			return "系统出现错误"
		}
		if(pay.userId == userID){
			return "1"   //这个人已经付钱了
		}
	}
	return "0"   //这里结束的话，就代表这个人没有付钱
}

//利用用户的id获取这个id的积分数
func GEtintegral(userId string)string{
	var intergral string
	rows,err:=db.Query("select integral from users where userid=$1",userId)

	if err!=nil{
		fmt.Println("用户积分查找失败1！",err)
		return "系统出现错误"
	}

	for rows.Next(){
		err = rows.Scan(&intergral)
		if(err!=nil){
			fmt.Println("用户积分查找失败2")
			return "系统出现错误"
		}
	}

	return intergral
}

func Getusername(userId string)string{
	var username string
	rows,err:=db.Query("select username from users where userid=$1",userId)

	if err!=nil{
		fmt.Println("查找用户名失败！",err)
		return "系统出现错误"
	}

	for rows.Next(){
		err = rows.Scan(&username)
		if(err!=nil){
			fmt.Println("查找用户名失败2",err)
			return "系统出现错误"
		}
	}
	return username
}



//执行收藏和取消收藏的指令
func Implement(contenTitle string,instruction string,titleId int,languageType string,userId int,collectiontime string)string{
	//上面代表的内容有：文章的title，执行收藏还是取消看收藏的指令，文章的id，语言类型，用户id
	// var collect collectioncase

	//注意这是处理事件
	if(instruction=="collect"){
		//收藏
		if(languageType=="日"){
			languageType = "japancase"
		}else{
			languageType = "koreacase"
		}
		stmt,err:=db.Prepare("insert into collection(userid,collectiontype,collectioncontentid,collectiontime) values($1,$2,$3,$4)")
		if err!=nil{
			fmt.Println("添加数据的时候出现的错误1",err)
			return "系统出现错误"
		}
		rs,err:=stmt.Exec(userId,languageType,titleId,collectiontime)

		if err!=nil{
			fmt.Println("添加数据的时候出现错误2",err)
		}
		fmt.Println("添加数据成功",rs)

	}
	if(instruction=="cancle"){
		//取消收藏
		if(languageType=="日"){
			languageType = "japancase"
		}else{
			languageType = "koreacase"
		}
		stmt,err:=db.Prepare("delete from collection where userid=$1 and collectiontype=$2 and collectioncontentid =$3")
		if err!=nil{
			fmt.Println("删除数据的时候出现的错误1",err)
			return "系统出现错误"
		}
		rs,err:=stmt.Exec(userId,languageType,titleId)

		if err!=nil{
			fmt.Println("删除数据的时候出现错误2",err)
			return "系统出现错误"
		}
		fmt.Println("删除数据成功",rs)
	}
	return "成功"
}

//收藏的初状态
func Statecollect(contenTitle string,titleId int,languageType string,userId int)string{
	var collect collection
	if(languageType=="日"){
		languageType = "japancase"
	}else{
		languageType = "koreacase"
	}

	//使用userid对文章进行判断
	rows,err:=db.Query("select * from collection where userid=$1",userId)

	if err!=nil{
		fmt.Println("查找收藏的状态的时候出现的错误1",err)
		return "系统出现错误"
	}
	
	for rows.Next(){
		err:=rows.Scan(&collect.userid,&collect.collectionid,&collect.collectiontype,&collect.collectioncontentid,&collect.collectiontime)
		if err!=nil{
			fmt.Println("查找收藏状态的时候出现错误2",err)
		}
		if(collect.collectiontype==languageType&&collect.collectioncontentid==titleId){
			return "collect"
		}
	}
	return "uncollect"
}


//数据库端的扣钱函数
func Pay(titleId string,userId string,integral string)string{
	
	//首先解决付钱的问题
	var all_integral int
	var need_integral int
	need_integral,_ = strconv.Atoi(integral)    //获取需要的积分
	fmt.Println(need_integral)
	rows,err:=db.Query("select integral from users where userid=$1",userId)
	if err!=nil{
		fmt.Println("查找用户名字相对应的积分的状态的时候出现的错误1",err)
		return "系统出现错误"
	}

	for rows.Next(){
		err:=rows.Scan(&all_integral)
		if err!= nil{
			fmt.Println("查找用户名字相对应的积分的状态的时候出现的错误12",err)
			return "系统出现错误"
		}

		fmt.Println(all_integral)

		if(all_integral == 0 || (all_integral<need_integral)){
			fmt.Println("积分不够")
			return "积分不够"
		}else{
			all_integral = all_integral - need_integral
			//把积分这钱保存在数据库中
			integral_data :=Saveintegral(all_integral,userId)
			if(integral_data=="保存成功"){
				pay_data:=SavePayData(titleId,userId)
				return pay_data
			}
		}
	}

	return "系统出现错误"
}

//保存积分函数
func Saveintegral(integral int,userId string) string{
	//是更新数据
	stmt, err := db.Prepare("update users set integral=$1 where userid=$2")

	if err != nil {
		//err
		fmt.Println("保存积分的时候出现错误1")
		return "系统出现错误"
	}

	_, err = stmt.Exec(integral,userId)
	if err != nil {
		fmt.Println("保存积分的时候出现错误2")
		return "系统出现错误"
	}

	return "保存成功"
}

//当保存成功之后，就将数据保存在已经付钱的数据库中
func SavePayData(titleId string,userId string)string{
	stmt, err := db.Prepare("insert into payment(useid, pointid) values($1,$2)")

	if err != nil {
		//err
		fmt.Println("保存付钱资料的时候出现错误1",err)
		return "系统出现错误"
	}

	_, err = stmt.Exec(userId,titleId)
	if err != nil {
		fmt.Println("保存付钱资料的时候出现错误2")
		return "系统出现错误"
	}

	return "保存成功"
}

