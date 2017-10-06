package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	repos "../lib/repos"
)

// HeroController that has methods as Request Handlers
type HeroController struct {
}

//FindAndDumpHeroes  dumps all Dota2 Heroes into our db
func (h *HeroController) FindAndDumpHeroes() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://api.opendota.com/api/heroes")
		if err != nil {
			// handle error
			panic(err)
		}
		defer resp.Body.Close()
		io.Copy(w, resp.Body)

	}
}

//GetAllHeroes gets all Heroes from db
func (h *HeroController) GetAllHeroes() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("yolo")

		heroes := repos.FindAll()
		json.NewEncoder(w).Encode(heroes)
	}
}

//InsertHero inserts Hero in db
// func (HeroController) InsertHero() func(w http.ResponseWriter, r *http.Request) {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 	}
// }
