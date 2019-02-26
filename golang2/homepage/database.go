package homepage

import(
	"io/ioutil"
	"fmt"
	_"github.com/lib/pq"
	"database/sql"
	"log"
)

var Db *sql.DB	

func init(){
	//the way of connect to database 
	var pgkey string = "postgres://postgres:password@129.204.193.192/homepage?sslmode=disable"
	var err error
	Db,err = sql.Open("postgres",pgkey)
	if err != nil {
		log.Fatal("sql open err :",err)
		return
	}
	err = Db.Ping()
	if err != nil {
		log.Fatal("sql ping err :",err)
	}
}

//Get the lastest fith news data 
func GetHomePageHotnewDate()(date [5]HomePageNews){
	command := "select imgurl, linkurl, title from homepagenews order by id desc limit 5"
	rows,err := Db.Query(command)
	defer rows.Close()
	if err!= nil{
		log.Println("error at gethomepagenewdate: ",err)
		return
	}
	index := 0
	for rows.Next() {
		rows.Scan(&date[index].Img_url, &date[index].Link_url, &date[index].Title)
		index++ 
	}
	return 
}

func GetHPADate(index int64)(date []ArticlaBox){
	command := "select imgurl, linkurl, brif, date from homepageartical order by id asc offset $1 limit 10"
	rows,err := Db.Query(command,index)
	var temp ArticlaBox
	defer rows.Close()
	if err!=nil{
		fmt.Println("error at GetHomePageArtical() , err:",err)
		return
	}
	for i := 0;rows.Next();i++ {
		rows.Scan(&temp.Img_url, &temp.Link_url, &temp.Brif, &temp.Date)
		date = append(date,temp)
	}
	return
}

//return images's byte that user needed
func Images(tag string, name string)[]byte{
	var  img_path string = "/home/ubuntu/DockerWorkPlace/Golang/source/images/"
	filepath := img_path + tag + name
	temp,err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Images() err: ", err)
	}
	return temp
}
