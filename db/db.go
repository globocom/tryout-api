package db

import mgo "gopkg.in/mgo.v2"

var Db *mgo.Database
var DbSession *mgo.Session

func DialDB() {
	DbSession, _ = mgo.Dial("localhost:27017")
	Db = DbSession.DB("tryout")
}

func Coll(collection string) *mgo.Collection {
	return Db.C(collection)
}
