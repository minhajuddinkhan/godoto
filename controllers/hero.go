package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	models "github.com/minhajuddinkhan/godoto/lib/models"
	repos "github.com/minhajuddinkhan/godoto/lib/repos"
)

var (
	heroRepo = repos.HeroRepo{}
)

// HeroController that has methods as Request Handlers
type HeroController struct {
}

//FindAndDumpHeroes  dumps all Dota2 Heroes into our db
func (h *HeroController) FindAndDumpHeroes() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		secret := r.Header.Get("secret")
		if secret != "godumpheroesforme" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(http.StatusText(http.StatusUnauthorized))
			return
		}

		resp, _ := http.Get("https://api.opendota.com/api/heroes")

		var heroes []models.Hero
		err := json.NewDecoder(resp.Body).Decode(&heroes)
		if err != nil {
			fmt.Println(err)
		}

		for _, hero := range heroes {
			heroRepo.InsertHero(&hero)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(heroes)

	}
}

//GetAllHeroes gets all Heroes from db
func (h *HeroController) GetAllHeroes() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		heroes := heroRepo.FindAll()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(heroes)
	}
}

//InsertHero inserts Hero in db
func (h *HeroController) InsertHero() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		var hero models.Hero
		json.NewDecoder(r.Body).Decode(&hero)
		fmt.Println(hero)
		err := heroRepo.InsertHero(&hero)
		if err != nil {
			fmt.Println("err executing insert query", err)
		}
		json.NewEncoder(w).Encode(hero)
	}

}

//GetHero gets hero by Id
func (h *HeroController) GetHero() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		heroID, _ := strconv.ParseInt(mux.Vars(r)["heroId"], 4, 10)
		hero := models.Hero{}
		err := heroRepo.FindOne(heroID, &hero)
		if err != nil {
			fmt.Println("err executing find query", err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hero)

	}
}
