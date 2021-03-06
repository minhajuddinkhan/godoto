package repos

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	mongo "github.com/minhajuddinkhan/godoto/lib/db"
	models "github.com/minhajuddinkhan/godoto/lib/models"
	"gopkg.in/mgo.v2"
)

var (
	heroes []models.Hero
)

type HeroRepo struct {
}

//FindAll gets all heroes
func (h *HeroRepo) FindAll() []models.Hero {

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
func (h *HeroRepo) DumpHeroes(heroes *[]models.Hero) error {

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
func (h *HeroRepo) InsertHero(hero *models.Hero) error {
	query := func(c *mgo.Collection) error {
		err := c.Insert(hero)
		if err != nil {
			fmt.Println("Cant insert", err)
		}
		return err
	}
	err := mongo.WithCollection("heroes", query)
	if err != nil {
		fmt.Println("error in insert", err)
	}
	return err
}

//FindOne finds one hero
func (h *HeroRepo) FindOne(heroID int64, hero *models.Hero) error {

	query := func(c *mgo.Collection) error {
		err := c.Find(bson.M{"id": heroID}).One(&hero)
		if err != nil {
			fmt.Println("query broke", err)

		}
		return err
	}

	err := mongo.WithCollection("heroes", query)
	if err != nil {
		fmt.Println("can't find one -repo layer")
	}
	return err

}
