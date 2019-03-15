package psql

import (
	"fmt"
)

type Feedback struct {
	Userid          float64 `json:"userid"`
	Feedbackid      string  `json:"feedbackid"`
	Feedbackcontent string  `json:"feedbackcontent"`
	Feedbacktype    string  `json:"feedbacktype"`
	Feedbacktime    string  `json:"feedbacktime"`
	Feedbackreplie  string  `json:"feedbackreplie"`
}

func Addfeedback(userid float64, feedbacktype string, feedbackcontent string) {

	stmt, err := db.Prepare("insert into feedback(userid,feedbacktype,feedbackcontent,feedbackreplie) values($1,$2,$3,$4)")
	CheckErr(err)
	//################################################## checkErr -> panic ?
	_, err = stmt.Exec(userid, feedbacktype, feedbackcontent, "无")
	CheckErr(err)
	fmt.Println("insert into feedback success")

}

func Userfeedback(tuserid float64) []Feedback {

	var (
		feedbacklist []Feedback
		feedback     Feedback
	)

	rows, err := db.Query("select * from feedback order by feedbackid desc")
	CheckErr(err)

	for rows.Next() {

		err = rows.Scan(&feedback.Feedbackid, &feedback.Userid, &feedback.Feedbackcontent, &feedback.Feedbacktype, &feedback.Feedbacktime, &feedback.Feedbackreplie)
		CheckErr(err)

		if tuserid == feedback.Userid {

			//fmt.Println(feedbacktime)

			feedbacklist = append(feedbacklist, feedback)
		}

	}

	return feedbacklist

}
