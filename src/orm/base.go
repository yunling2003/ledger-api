package orm

import (
	"ledger-api/src/config"

	"log"

	"gopkg.in/mgo.v2"
)

var (
	db *mgo.Database
)

func init() {
	mongoURL := config.All["mongourl"]
	log.Println("mongo url is: " + mongoURL)
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db = session.DB("ledger")
}
