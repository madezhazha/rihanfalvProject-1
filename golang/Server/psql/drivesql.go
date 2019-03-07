package psql

import(
	// _"github.com/lib/pq"
	"fmt"
	"strings"
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



type concretelaw1 struct{
	ID string 
	viewpoint string
	casetitle string
	header string
	Type string
}

type collectioncase struct{
	ID string
	collectiontype string
	collectiontitle string
	collectiontitleid string
	userId string
}




//获取所有的数据
func Getalldata(languageType string)interface{}{

	all_data := make(map[string]map[string]string)

	//数据库的查询
	rows,err := db.Query("select * from casething where type=$1",languageType)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event,&info.Type)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误2"
		}

		//将数据放在了一块
		all_data[info.ID] = make(map[string]string)
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

func Getfirstfloor(data string,languageType string)interface{}{
	all_data := make(map[string]map[string]string)

	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from casething where causeofaction=$1 and type=$2",data,languageType)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event,&info.Type)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误2"
		}

		//将数据放在了一块
		all_data[info.ID] = make(map[string]string)
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


func Getreason(data string,languageType string)interface{}{
	all_data := make(map[string]map[string]string)

	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from casething where concretecasetype=$1 and type=$2",data,languageType)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event,&info.Type)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误2"
		}

		//将数据放在了一块
		all_data[info.ID] = make(map[string]string)
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


func Getlevel(data string,languageType string)interface{}{
	all_data := make(map[string]map[string]string)

	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from casething where trialgrade=$1 and type=$2",data,languageType)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event,&info.Type)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误2"
		}

		//将数据放在了一块
		all_data[info.ID] = make(map[string]string)
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

func Getsecondfloor(data string,languageType string)interface{}{
	all_data := make(map[string]map[string]string)


	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from casething where legalprinciple=$1 and type=$2",data,languageType)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event,&info.Type)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误2"
		}

		//将数据放在了一块
		all_data[info.ID] = make(map[string]string)
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

func Gettime(data string,languageType string)interface{}{
	all_data := make(map[string]map[string]string)

	a:=0

	rows,err := db.Query("select * from casething where type=$1",languageType)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event,&info.Type)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误2"
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


//法官观点
func Gettext(data string)interface{}{
	var con concretelaw1

	//创建一个单个集合的承接数据
	all:=make(map[string]string)
	rows,err:=db.Query("select * from point where casetitle=$1",data) 

	if err!=nil{
		fmt.Println(err)
		return "系统出席那错误"
	}
	for rows.Next(){
		err:=rows.Scan(&con.ID,&con.viewpoint,&con.casetitle,&con.header,&con.Type)
		if err!=nil{
			fmt.Println(err)
			return "系统出现错误"
		}
		all["viewpoint"]=con.viewpoint
		all["casetitle"]=con.casetitle
		all["header"]=con.header
		all["type"]=con.Type
		all["ID"]=con.ID
	}

	defer rows.Close()

	return all
}


//执行收藏和取消收藏的指令
func Implement(contenTitle string,instruction string,titleId string,languageType string,userId string){
	// var collect collectioncase

	//注意这是处理事件
	if(instruction=="collect"){
		//收藏
		stmt,err:=db.Prepare("insert into collectioncase(userid,collectiontype,collectiontitle,collectiontitleid) values($1,$2,$3,$4)")
		if err!=nil{
			fmt.Println("添加数据的时候出现的错误1",err)
			return 
		}
		rs,err:=stmt.Exec(userId,languageType,contenTitle,titleId)

		if err!=nil{
			fmt.Println("添加数据的时候出现错误2",err)
		}
		fmt.Println("添加数据成功",rs)

	}
	if(instruction=="cancle"){
		//取消收藏
		stmt,err:=db.Prepare("delete from collectioncase where userid=$1 and collectiontype=$2 and collectiontitle =$3 and collectiontitleid=$4")
		if err!=nil{
			fmt.Println("删除数据的时候出现的错误1",err)
			return 
		}
		rs,err:=stmt.Exec(userId,languageType,contenTitle,titleId)

		if err!=nil{
			fmt.Println("删除数据的时候出现错误2",err)
			return 
		}
		fmt.Println("删除数据成功",rs)
	}
}

//收藏的初状态
func Statecollect(contenTitle string,titleId string,languageType string,userId string)string{
	var collect collectioncase

	//使用userid对文章进行判断
	rows,err:=db.Query("select * from collectioncase where userid=$1",userId)

	if err!=nil{
		fmt.Println("查找收藏的状态的时候出现的错误1",err)
		return "错误"
	}
	
	for rows.Next(){
		err:=rows.Scan(&collect.ID,&collect.userId,&collect.collectiontype,&collect.collectiontitle,&collect.collectiontitleid)
		if err!=nil{
			fmt.Println("查找收藏状态的时候出现错误2",err)
		}
		if(collect.collectiontype==languageType&&collect.collectiontitle==contenTitle&&collect.collectiontitleid==titleId){
			return "collect"
		}
	}
	return "uncollect"
}



