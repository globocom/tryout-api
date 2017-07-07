package db

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var Db *mgo.Database
var DbSession *mgo.Session

func DialDB() error {
	url := "localhost:27017"
	envUrl := os.Getenv("DATABASE_URL")
	if envUrl != "" {
		url = envUrl
	}
	dialInfo, err := mgo.ParseURL(url)
	fmt.Printf("%#v", dialInfo)
	dialInfo.FailFast = true
	DbSession, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return err
	}
	fmt.Print("xxx")
	Db = DbSession.DB("tryout_mongo")
	return nil
}

func Coll(collection string) *mgo.Collection {
	return Db.C(collection)
}
