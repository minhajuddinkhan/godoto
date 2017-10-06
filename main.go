package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/godoto/web"
)

func main() {

	r := mux.NewRouter()
	r = web.Up(*r)

	http.ListenAndServe(":3000", r)
}
