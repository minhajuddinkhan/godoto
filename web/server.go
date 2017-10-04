package web

import (
	"github.com/gorilla/mux"
	"../controllers"
)

func Up(r mux.Router)  *mux.Router {

	r.HandleFunc("/health", controllers.Health);

	return &r;


}