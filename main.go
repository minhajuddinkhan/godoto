package main

import (
	"github.com/gorilla/mux"
	"./web"
	"net/http"
)

func main() {
	
	r := mux.NewRouter();
	r = web.Up(*r);

	http.ListenAndServe(":3000", r)
}
