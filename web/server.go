package web

import (
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/godoto/controllers"
)

var (
	heroCtrl = controllers.HeroController{}
)

func Up(r mux.Router) *mux.Router {

	r.HandleFunc("/heroes", heroCtrl.FindAndDumpHeroes()).Methods("POST")
	r.HandleFunc("/heroes", heroCtrl.GetAllHeroes()).Methods("GET")
	return &r
}
