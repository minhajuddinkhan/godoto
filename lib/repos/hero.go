package repos

import (
	mongo "../db"
	models "../models"
	"gopkg.in/mgo.v2"
)

var (
	heroes []models.Hero
)

//FindAll gets all heroes
func FindAll() []models.Hero {

	heroes = []models.Hero{}
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&heroes)
	}

	err := mongo.WithCollection("heroes", query)
	if err != nil {
		panic(err)
	}
	return heroes
}

// DumpHeroes dumps all heroes in monogdb
func DumpHeroes(heroes []models.Hero) {

	heroes = []models.Hero{}
	query := func(c *mgo.Collection) error {
		return c.Insert(heroes)
	}
	err := mongo.WithCollection("heroes", query)
	if err != nil {
		panic(err)
	}
	return
}
