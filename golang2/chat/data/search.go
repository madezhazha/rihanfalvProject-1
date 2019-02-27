package data

import "fmt"

// RsByCondition 根据查询条件查询出结果
func RsByCondition(condition []string) (thread []Thread, err error) {
	for i, value := range condition {
		rows, err := Db.Query("SELECT topicid,userid,topictitle,topiccontent,creationtime FROM topics where tag like $1 ORDER BY topicid", "%"+value+"%")
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			conv := Thread{}
			if err = rows.Scan(&conv.ID, &conv.Userid, &conv.Topictitle, &conv.Topiccontent, &conv.Creationtime); err != nil {
				return nil, err
			}
			if i == 0 {
				thread = append(thread, conv)
			} else {
				flag := true
				for _, value := range thread {
					if value.ID == conv.ID {
						flag = false
					}
				}
				if flag == true {
					thread = append(thread, conv)
				}
			}
		}
		fmt.Println("--------------------")
		// fmt.Println(value)
		for _, value := range thread {
			fmt.Println(value.ID)
		}
	}
	return
}
