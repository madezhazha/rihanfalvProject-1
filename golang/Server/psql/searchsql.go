package psql

import (
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

var class string


type Searchbox struct{
	ID	int `json:"ID"`
	Type	string `json:"Type"`
	Title string `json:"Title"`
	Author string `json:"Author"`
	Time string `json:"Date"`
	Content string `json:"Content"`
	Length int `json:"Length"`
	Label string `json:"Label"`
	Labelbox []string `json:"Labelbox"`
	Value int `json:"Value"`
	Classify string `json:"Classify"`
}

func Getclass(readkey string,readcountry string,readclass string,searchlist []Searchbox)[]Searchbox{
	if readcountry=="日"{
		readcountry="japan"
	}
	if readcountry=="韩"{
		readcountry="korea"
	}
	if readclass=="全部"{
		class=readcountry+"legal"
		searchlist=Legalsearch(readkey,searchlist)
		class=readcountry+"thesis"
		searchlist=Thesissearch(readkey,searchlist)

	}
	if readclass=="法律条文"{
		class=readcountry+"legal"
		searchlist=Legalsearch(readkey,searchlist)
	}
	if readclass=="案例"{	
		class=readcountry+"analysis"
	}
	if readclass=="论文"{	
		class=readcountry+"thesis"
		searchlist=Thesissearch(readkey,searchlist)

	}
	return searchlist
}

func Legalsearch(readkey string,searchlist []Searchbox)[]Searchbox{//从法律数据库查询
	var m Searchbox
		 //得到查找的语句，%_%表示前后模糊查找
		searchstr := "select * from "+class+" where LegalTitle like '%" + readkey + "%' or LegalContent like '%" + readkey + "%' " 
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
func Thesissearch(readkey string,searchlist []Searchbox)[]Searchbox{//从论文数据库查询
	var m Searchbox
	 //得到查找的语句，%_%表示前后模糊查找
		searchstr := "select * from "+class+" where ThesisTitle like '%" + readkey + "%' or ThesisContent like '%" + readkey + "%' " 
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

func Scoreofsearch(searchlist []Searchbox,readkey string){//判断内容的相关度
	i:=0;
	lenth:=len(searchlist)
	for ;i<lenth;i++{
        searchlist[i].Value=0
		titlecount:=strings.Count(searchlist[i].Title, readkey)//获取内容中key的出现次数
		searchlist[i].Value=searchlist[i].Value+titlecount*100
		//fmt.Println("标题包含：",titlecount) 
		contentcount:=strings.Count(searchlist[i].Content, readkey)
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



