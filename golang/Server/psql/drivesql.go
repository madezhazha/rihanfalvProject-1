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




//获取所有的数据
func Getalldata()interface{}{

	all_data := make(map[string]map[string]string)

	//数据库的查询
	rows,err := db.Query("select * from casething")

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

func Getfirstfloor(data string)interface{}{
	all_data := make(map[string]map[string]string)

	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from casething where causeofaction=$1",data)

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


func Getreason(data string)interface{}{
	all_data := make(map[string]map[string]string)

	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from casething where concretecasetype=$1",data)

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


func Getlevel(data string)interface{}{
	all_data := make(map[string]map[string]string)

	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from casething where trialgrade=$1",data)

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

func Getsecondfloor(data string)interface{}{
	all_data := make(map[string]map[string]string)


	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from casething where legalprinciple=$1",data)

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

func Gettime(data string)interface{}{
	all_data := make(map[string]map[string]string)

	a:=0

	rows,err := db.Query("select * from casething")

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
	}

	defer rows.Close()

	return all
}
