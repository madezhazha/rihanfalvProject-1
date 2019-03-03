package psql

import (
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

var class[] string


type Searchbox struct{
	ID	int `json:"ID"`
	Type	string `json:"Type"`
	Title string `json:"Title"`
	Author string `json:"Author"`
	Time string `json:"Date"`
	Content string `json:"Content"`
	Length int `json:"Length"`
	Label string `json:"Label"`
	Value int `json:"Value"`
}

func Getclass(readkey string,readclass string,searchlist []Searchbox)[]Searchbox{
	if readclass=="全部"{
		class=[]string{"japanlegal","korealegal"}
		searchlist=Legalsearch(readkey,searchlist)
		class=[]string{"japanthesis","koreathesis"}
		searchlist=Thesissearch(readkey,searchlist)
	}
	if readclass=="法律条文"{
		class=[]string{"japanlegal","korealegal"}
		searchlist=Legalsearch(readkey,searchlist)
	}
	if readclass=="案例"{		
		class=[]string{"japananalysis","koreaanalysis"}

	}
	if readclass=="论文"{	
		class=[]string{"japanthesis","koreathesis"}
		searchlist=Thesissearch(readkey,searchlist)

	}
	return searchlist
}

func Legalsearch(readkey string,searchlist []Searchbox)[]Searchbox{//从数据库查询
	var m Searchbox
	for  i:=0;i<len(class);i++{		 //得到查找的语句，%_%表示前后模糊查找
		searchstr := "select * from "+class[i]+" where LegalTitle like '%" + readkey + "%' or LegalContent like '%" + readkey + "%' " 
		rows, err := db.Query(searchstr)
		if err != nil {
			fmt.Println("ERROR:", err)
			return searchlist
		} //检查错误
		for rows.Next() { //将rows赋值
		rows.Scan(&m.ID, &m.Type, &m.Title, &m.Content, &m.Label)
		searchlist=append(searchlist,m)
		}	
	}
	return searchlist
}
func Thesissearch(readkey string,searchlist []Searchbox)[]Searchbox{//从数据库查询
	var m Searchbox
	for  i:=0;i<len(class);i++{		 //得到查找的语句，%_%表示前后模糊查找
		searchstr := "select * from "+class[i]+" where ThesisTitle like '%" + readkey + "%' or ThesisContent like '%" + readkey + "%' " 
		rows, err := db.Query(searchstr)
		if err != nil {
			fmt.Println("ERROR:", err)
			return searchlist
		} //检查错误
		for rows.Next() { //将rows赋值
		rows.Scan(&m.ID, &m.Title, &m.Author, &m.Time, &m.Content,&m.Length,&m.Label)
		searchlist=append(searchlist,m)
		fmt.Println(m)
		}	
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
		fmt.Println("标题包含：",titlecount) 
		contentcount:=strings.Count(searchlist[i].Content, readkey)
		searchlist[i].Value=searchlist[i].Value+contentcount*20
		 fmt.Println("内容包含",contentcount)  
		 fmt.Println("相关度：",searchlist[i].Value)
		 //分数=标题出现次数*100+内容出现次数*20(+点击次数)
	}
}	
func SelectSort(searchlist []Searchbox) { //冒泡排序
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



