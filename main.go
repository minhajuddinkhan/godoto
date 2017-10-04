package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	clientName string
	Phone      string
}

type Person struct {
	Name  string
	Phone string
}

func main() {

	session, err := mgo.Dial("mongodb://localhost:27017/test")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.

	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("test").C("clients")

	client := Client{}
	colErr := collection.Find(bson.M{}).One(&client)
	if colErr != nil {
		fmt.Println(colErr)
	}
	fmt.Println(client)
}
