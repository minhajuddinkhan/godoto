package mongo

import (
	"gopkg.in/mgo.v2"
)

const (
	db       string = "test"
	mongoURI string = "mongodb://localhost:27017" + db
)

var (
	mgoSession *mgo.Session
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(mongoURI)
		if err == nil {
			panic(err)
		}

	}
	return mgoSession.Clone()
}

func WithCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(db).C(collection)
	return s(c)
}
