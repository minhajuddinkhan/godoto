package repos

import (
	"fmt"

	mongo "github.com/minhajuddinkhan/godoto/lib/db"
	models "github.com/minhajuddinkhan/godoto/lib/models"
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
		fmt.Errorf("error in query", err)
	}
	return heroes
}

// DumpHeroes dumps all heroes in monogdb
func DumpHeroes(heroes *[]models.Hero) error {

	query := func(c *mgo.Collection) error {
		return c.Insert(heroes)
	}
	err := mongo.WithCollection("heroes", query)
	if err != nil {
		fmt.Println("query broke at repo layer")
	} else {
		fmt.Println("query ran fine at repo layer dump.")
	}
	return err
}

//InsertHero repo layer
func InsertHero(hero *models.Hero) error {
	query := func(c *mgo.Collection) error {
		return c.Insert(&hero)
	}
	return mongo.WithCollection("heroes", query)
}
