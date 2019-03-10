package psql

import (
	"fmt"
)

type CollectionsMsg struct {
	CollectionTime      string
	CollectionTitle     string
	CollectionID        int
	CollectionContentID int
	CollectionContent   string
	CollectionType      string
}

func SqlSelect(UserID int) []CollectionsMsg {

	//var MCollectionsMsg [50]CollectionsMsg

	//var NullCollection [50]CollectionsMsg

	//MCollectionsMsg = NullCollection

	var (
		CollectionsMsgList []CollectionsMsg
		CollectionsMsg     CollectionsMsg
	)

	topiceRows, err := db.Query("SELECT collectiontime,topictitle,collectionid,collectioncontentid,topiccontent,collectiontype FROM collection join topics on(collection.collectioncontentid = topics.topicid)    where userid=$1 and collectiontype='topic' ", UserID)
	checkErr(err)

	for topiceRows.Next() {

		err = topiceRows.Scan(&CollectionsMsg.CollectionTime, &CollectionsMsg.CollectionTitle, &CollectionsMsg.CollectionID, &CollectionsMsg.CollectionContentID, &CollectionsMsg.CollectionContent, &CollectionsMsg.CollectionType)
		checkErr(err)

		CollectionsMsgList = append(CollectionsMsgList, CollectionsMsg)
	}

	caseRows, err := db.Query("SELECT collectiontime,casetitle,collectionid,collectioncontentid,casecontent,collectiontype FROM collection join casething on(collection.collectioncontentid = casething.id)    where userid=$1 and collectiontype='koreacase' ", UserID)
	checkErr(err)

	for caseRows.Next() {

		err = caseRows.Scan(&CollectionsMsg.CollectionTime, &CollectionsMsg.CollectionTitle, &CollectionsMsg.CollectionID, &CollectionsMsg.CollectionContentID, &CollectionsMsg.CollectionContent, &CollectionsMsg.CollectionType)
		checkErr(err)

		CollectionsMsgList = append(CollectionsMsgList, CollectionsMsg)
	}

	japanthesisRows, err := db.Query("SELECT collectiontime,thesistitle,collectionid,collectioncontentid,thesiscontent,collectiontype FROM collection join japanthesis on(collection.collectioncontentid = japanthesis.thesisid)    where userid=$1 and collectiontype='japanthesis' ", UserID)
	checkErr(err)

	for japanthesisRows.Next() {

		err = japanthesisRows.Scan(&CollectionsMsg.CollectionTime, &CollectionsMsg.CollectionTitle, &CollectionsMsg.CollectionID, &CollectionsMsg.CollectionContentID, &CollectionsMsg.CollectionContent, &CollectionsMsg.CollectionType)
		checkErr(err)

		CollectionsMsgList = append(CollectionsMsgList, CollectionsMsg)
	}
	koreathesisRows, err := db.Query("SELECT collectiontime,thesistitle,collectionid,collectioncontentid,thesiscontent,collectiontype FROM collection join koreathesis on(collection.collectioncontentid = koreathesis.thesisid)    where userid=$1 and collectiontype='koreathesis' ", UserID)
	checkErr(err)

	for koreathesisRows.Next() {

		err = koreathesisRows.Scan(&CollectionsMsg.CollectionTime, &CollectionsMsg.CollectionTitle, &CollectionsMsg.CollectionID, &CollectionsMsg.CollectionContentID, &CollectionsMsg.CollectionContent, &CollectionsMsg.CollectionType)
		checkErr(err)

		CollectionsMsgList = append(CollectionsMsgList, CollectionsMsg)
	}

	return CollectionsMsgList
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
