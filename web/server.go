package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/godoto/controllers"
)

var (
	heroCtrl = controllers.HeroController{}
)

func Up() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/heroes", heroCtrl.InsertHero()).Methods("POST")
	r.HandleFunc("/heroes", heroCtrl.FindAndDumpHeroes()).Methods("PUT")
	r.HandleFunc("/heroes", heroCtrl.GetAllHeroes()).Methods("GET")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/", r)
}
