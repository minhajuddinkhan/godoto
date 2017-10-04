package web

import (
	"github.com/gorilla/mux"
	"../controllers"
)

func Up(r mux.Router)  *mux.Router {

	r.HandleFunc("/heroes", controllers.Hero);

	return &r;


}