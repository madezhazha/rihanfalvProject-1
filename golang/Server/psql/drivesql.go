package psql

import(
	//"database/sql"
	//_"github.com/lib/pq"
	"fmt"
	//"log"
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
}


type concretelaw struct{
	ID string
	judgepoint string
	firstinstance string
	secondtrial string
	thirdtrial string
	title string
	publicoffice string
	plaintiff string
	agent string
	defendant string
	counsel string
	trialgrade string
	firstcourt string
	firstpeople string
	secondcourt string
	secondpeople string
	retrial string
	retrialpeople string
	firsttime string
	secondtime string
	retrialtime string
}



//获取所有的数据
func Getalldata()interface{}{

	all_data := make(map[string]map[string]string)


	//数据库的查询
	rows,err := db.Query("select * from content")

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event)

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

	}

	rows.Close()

	return all_data
}

func Getfirstfloor(data string)interface{}{
	all_data := make(map[string]map[string]string)

	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from content where causeofaction=$1",data)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event)

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
	}

	rows.Close()

	return all_data
}


func Getreason(data string)interface{}{
	all_data := make(map[string]map[string]string)

	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from content where concretecasetype=$1",data)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event)

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
	}

	rows.Close()

	return all_data
}


func Getlevel(data string)interface{}{
	all_data := make(map[string]map[string]string)

	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from content where trialgrade=$1",data)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event)

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
	}

	rows.Close()

	return all_data
} 

func Getsecondfloor(data string)interface{}{
	all_data := make(map[string]map[string]string)


	//格式化输出
	// first_cmd := fmt.Sprintf("select * from content where causetype=%s",data)
	//数据库的查询
	rows,err := db.Query("select * from content where legalprinciple=$1",data)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event)

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
	}

	rows.Close()

	return all_data
}

func Gettime(data string)interface{}{
	all_data := make(map[string]map[string]string)

	a:=0

	rows,err := db.Query("select * from content")

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}

	var info Information

	for rows.Next(){
		err:=rows.Scan(&info.ID,&info.trialtime,&info.causetype,&info.concretecasetype,&info.judgetype,&info.trialgrade,&info.title,&info.event)

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

			a=0;
		}
	}

	rows.Close()

	return all_data

}


func Gettext(data string)interface{}{
	var con concretelaw
	var all_data string

	//创建一个单个集合承接数据
	all := make(map[string]string)

	rows,err := db.Query("select * from content where title=$1",data)

	if err!=nil{
		fmt.Println(err)
		return "系统出现错误"
	}


	for rows.Next(){
		err:=rows.Scan(&con.ID,&con.judgepoint,&con.firstinstance,&con.secondtrial,&con.thirdtrial,&con.title,&con.publicoffice,&con.plaintiff,&con.agent,&con.defendant,&con.counsel,&con.trialgrade,&con.firstcourt,&con.firstpeople,&con.secondcourt,&con.secondpeople,&con.retrial,&con.retrialpeople,&con.firsttime,&con.secondtime,&con.retrialtime)

		if err!=nil{
			fmt.Println(err)
			return "系统出现错误2"
		}



		all_data=con.judgepoint
		//数据的替换，可以将数据显示出来
		all_data=strings.Replace(all_data,"<","&lt;",-1)
		all_data=strings.Replace(all_data,">","&gt;",-1)
		all_data=strings.Replace(all_data,"\n","<br>",-1)
		all_data=strings.Replace(all_data,"\t","&nbsp;&nbsp;&nbsp;",-1)
		all_data=strings.Replace(all_data,"\n","<br>",-1)
		all_data=strings.Replace(all_data," ","&nbsp;",-1)

		all["judgepoint"]=all_data
		all["firstinstance"]=con.firstinstance
		all["secondtrial"]=con.secondtrial
		all["thirdtrial"]=con.thirdtrial
		all["title"]=con.title
		all["publicoffice"]=con.publicoffice
		all["plaintiff"]=con.plaintiff
		all["agent"]=con.agent
		all["defendant"]=con.defendant
		all["counsel"]=con.counsel
		all["trialgrade"]=con.trialgrade
		all["firstcourt"]=con.firstcourt
		all["firstpeople"]=con.firstpeople
		all["secondcourt"]=con.secondcourt
		all["secondpeople"]=con.secondpeople
		all["retrial"]=con.retrial
		all["retrialpeople"]=con.retrialpeople
		all["firsttime"]=con.firsttime
		all["secondtime"]=con.secondtime
		all["retrialtime"]=con.retrialtime
	}

	rows.Close()

	return all
}