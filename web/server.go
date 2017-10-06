package web

import (
	"../controllers"
	"github.com/gorilla/mux"
)

var (
	heroCtrl = controllers.HeroController{}
)

func Up(r mux.Router) *mux.Router {

	r.HandleFunc("/heroes", heroCtrl.InsertHero()).Methods("POST")
	r.HandleFunc("/heroes", heroCtrl.FindAndDumpHeroes()).Methods("PUT")
	r.HandleFunc("/heroes", heroCtrl.GetAllHeroes()).Methods("GET")
	return &r
}
