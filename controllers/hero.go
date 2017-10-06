package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/minhajuddinkhan/godoto/lib/models"
	repos "github.com/minhajuddinkhan/godoto/lib/repos"
)

// HeroController that has methods as Request Handlers
type HeroController struct {
}

//FindAndDumpHeroes  dumps all Dota2 Heroes into our db
func (h *HeroController) FindAndDumpHeroes() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		resp, _ := http.Get("https://api.opendota.com/api/heroes")

		heroes := []models.Hero{}
		err := json.NewDecoder(resp.Body).Decode(&heroes)
		if err != nil {
			fmt.Println(err)
		}
		repoErr := repos.DumpHeroes(heroes)
		fmt.Println(repoErr)
		json.NewEncoder(w).Encode(heroes)

	}
}

//GetAllHeroes gets all Heroes from db
func (h *HeroController) GetAllHeroes() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		heroes := repos.FindAll()
		json.NewEncoder(w).Encode(heroes)
	}
}

//InsertHero inserts Hero in db
func (h *HeroController) InsertHero() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		roles := []string{"key"}
		repos.InsertHero(&hero)

	}
}
