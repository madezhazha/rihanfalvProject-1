package psql
import(
	"io/ioutil"
	"fmt"
	"log"
)
type HomePageNews struct{
	Img_url	string 			`json:"imgurl"`
	Link_url string 		`json:"linkurl"`
	Title	string 			`json:"title"`
}
type ArticlaBox struct{
	Img_url	string 			`json:"imgurl"`
	Link_url string 		`json:"linkurl"`
	Brif	string 			`json:"brief"`
	Date	string 			`json:"date"`
}
//Get the lastest fith news data 
func GetHomePageHotnewDate()(date [5]HomePageNews){
	command := "select imgurl, linkurl, title from homepagenews order by id desc limit 5"
	rows,err := db.Query(command)
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

//return homepage aritical data
func GetHPADate(index int64)(date []ArticlaBox){
	command := "select imgurl, linkurl, brif, date from homepageartical order by id asc offset $1 limit 10"
	rows,err := db.Query(command,index)
	var temp ArticlaBox
	defer rows.Close()
	if err!=nil{
		fmt.Println("error at GetHomePageArtical() , err:",err)
		return
	}
	for i := 0;rows.Next();i++ {
		var str string
		rows.Scan(&temp.Img_url, &temp.Link_url, &temp.Brif, &str)
		temp.Date = (str[0:10])
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
