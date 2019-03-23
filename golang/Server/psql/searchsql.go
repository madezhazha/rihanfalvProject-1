package psql

import (
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

var class string


type Searchbox struct{
	ID	int `json:"ID"`//
	Type	string `json:"Type"`//法律类型、案例类型
	Title string `json:"Title"`//通用
	Author string `json:"Author"`//论文作者
	Time string `json:"Date"`//论文、案例时间
	Content string `json:"Content"`//各种内容，通用
	Length int `json:"Length"`//论文长度
	Label string `json:"Label"`//标签
	Labelbox []string `json:"Labelbox"`//标签box
	Value int `json:"Value"`//用来排序的参数
	Classify string `json:"Classify"`//区分法律、论文、案例
	Country string `json:"Country"`//案例区分国家，淼哥出来挨打
	Causeofaction string `json:"Causeofaction"`//案例参数1
	Concretecasetype string `json:"Concretecasetype"`//案例参数2
	Legalprinciple string `json:"Legalprinciple"`//案例参数3
	Trialgrade string `json:Trialgrade`//案例参数4
}

func Getclass(readkey []string,readcountry string,readclass string,readorder string,searchlist []Searchbox)[]Searchbox{
	if readcountry=="日"{
		readcountry="japan"
	}
	if readcountry=="韩"{
		readcountry="korea"
	}
	if readclass=="全部"{
		class=readcountry+"legal"
		searchlist=Legalsearch(readkey,readorder,searchlist)
		class=readcountry+"thesis"
		searchlist=Thesissearch(readkey,readorder,searchlist)
		searchlist=Analysissearch(readkey,readorder,searchlist)

	}
	if readclass=="法律条文"{
		class=readcountry+"legal"
		searchlist=Legalsearch(readkey,readorder,searchlist)
	}
	if readclass=="案例"{	
		searchlist=Analysissearch(readkey,readorder,searchlist)
	}
	if readclass=="论文"{	
		class=readcountry+"thesis"
		searchlist=Thesissearch(readkey,readorder,searchlist)

	}
	return searchlist
}

func Legalsearch(readkey []string,readorder string,searchlist []Searchbox)[]Searchbox{//从法律数据库查询
	var m Searchbox
	str1:=" LegalTitle like '%"+readkey[0]+"%' "
	str2:=" Legalcontent like '%"+readkey[0]+"%' "
	length:=len(readkey)
	for i:=1;i<length;i++{
		if readkey[i]==""{
			continue;
		}
	str1+=" and LegalTitle like '%"+readkey[i]+"%' "
	str2+=" and Legalcontent like '%"+readkey[i]+"%' "
	}
	searchstr := "select * from "+class+" where "+str1+ " or "+ str2 //拼接到查找的语句，%_%表示前后模糊查找，sql里and优先级＞or
	if readorder=="onlytitle"{
			searchstr = "select * from "+class+" where "+ str1
		} 
		fmt.Println(searchstr)
		rows, err := db.Query(searchstr)
		if err != nil {
			fmt.Println("ERROR:", err)
			return searchlist
		} //检查错误
		for rows.Next() { //将rows赋值
		rows.Scan(&m.ID, &m.Type, &m.Title, &m.Content, &m.Label)
		m.Labelbox=strings.Split(m.Label,"/")
		m.Classify="法律"
		searchlist=append(searchlist,m)
		}	
	return searchlist
}
func Thesissearch(readkey []string,readorder string,searchlist []Searchbox)[]Searchbox{//从论文数据库查询
	var m Searchbox
	 //得到查找的语句，%_%表示前后模糊查找
	 str1:=" ThesisTitle like '%"+readkey[0]+"%' "
	str2:=" ThesisContent like '%"+readkey[0]+"%' "
	length:=len(readkey)
	for i:=1;i<length;i++{
		if readkey[i]==""{
			continue;
		}
	str1+=" and ThesisTitle like '%"+readkey[i]+"%' "
	str2+=" and ThesisContent like '%"+readkey[i]+"%' "
	}
	searchstr := "select * from "+class+" where "+str1+ " or "+ str2 
		if readorder=="onlytitle"{
			searchstr = "select * from "+class+" where "+str1
		} 
		fmt.Println(searchstr)

		rows, err := db.Query(searchstr)
		if err != nil {
			fmt.Println("ERROR:", err)
			return searchlist
		} //检查错误
		for rows.Next() { //将rows赋值
		rows.Scan(&m.ID, &m.Title, &m.Author, &m.Time, &m.Content,&m.Length,&m.Label)
		m.Labelbox=strings.Split(m.Label,"/")

		
		m.Classify="论文"
		searchlist=append(searchlist,m)
		}	

	return searchlist
}

func Analysissearch(readkey []string,readorder string,searchlist []Searchbox)[]Searchbox{//从案例数据库查询
	var m Searchbox
	 //得到查找的语句，%_%表示前后模糊查找
	 str1:=" casetitle like '%"+readkey[0]+"%' "
	 str2:=" casecontent like '%"+readkey[0]+"%' "
	 length:=len(readkey)
	 for i:=1;i<length;i++{
		if readkey[i]==""{
			continue;
		}
	 str1+=" and casetitle like '%"+readkey[i]+"%' "
	 str2+=" and casecontent like '%"+readkey[i]+"%' "
	 }
	 searchstr := "select * from casething where "+str1+ " or "+ str2 
	 if readorder=="onlytitle"{
		 searchstr = "select * from casething where "+str1
	 } 
	 fmt.Println(searchstr)
		rows, err := db.Query(searchstr)
		if err != nil {
			fmt.Println("ERROR:", err)
			return searchlist
		} //检查错误
		for rows.Next() { //将rows赋值
		rows.Scan(&m.ID, &m.Time, &m.Causeofaction, &m.Concretecasetype, &m.Legalprinciple,&m.Trialgrade,&m.Title,&m.Content,&m.Country)
		m.Classify="案例"
		searchlist=append(searchlist,m)
		}	

	return searchlist
}

func Scoreofsearch(searchlist []Searchbox,readkey []string){//判断内容的相关度
	i:=0;
	lenth:=len(searchlist)
	for ;i<lenth;i++{
        searchlist[i].Value=0
		titlecount:=strings.Count(searchlist[i].Title, readkey[0])//获取内容中key的出现次数
		searchlist[i].Value=searchlist[i].Value+titlecount*1000
		//fmt.Println("标题包含：",titlecount) 
		contentcount:=strings.Count(searchlist[i].Content, readkey[0])
		searchlist[i].Value=searchlist[i].Value+contentcount*5
		// fmt.Println("内容包含",contentcount)  
		// fmt.Println("相关度：",searchlist[i].Value)
		 //分数=标题出现次数*100+内容出现次数*5(+点击次数)
	}
}	
func SelectSort(searchlist []Searchbox) { //排序
	length := len(searchlist) 
	fmt.Println("搜索到目标项：",length)
	for i := 0; i < length-1; i++ { 
		min := i //每一趟的开始把首元素的下标作为最小元素的下标 
		for j := i + 1; j < length; j++ {
			if searchlist[j].Value > searchlist[min].Value { 
				min = j } 
				} 
				if min != i {
					searchlist[i], searchlist[min] = searchlist[min], searchlist[i] } } }



