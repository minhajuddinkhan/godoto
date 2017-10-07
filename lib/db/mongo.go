package mongo

import (
	"gopkg.in/mgo.v2"
)

const (
	db       string = "godoto"
	mongoURI string = "mongodb://godoto:godoto@ds013495.mlab.com:13495/" + db
)

var (
	mgoSession *mgo.Session
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(mongoURI)
		if err != nil {
			panic(err)
		}

	}
	return mgoSession.Clone()
}

//WithCollection C wrapper
func WithCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(db).C(collection)
	return s(c)
}
