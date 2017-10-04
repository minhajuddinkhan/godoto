package main

import (
	"github.com/gorilla/mux"
	"./web"
	"net/http"
)

type Client struct {
	clientName string
	Phone      string
}

type Person struct {
	Name  string
	Phone string
}

func main() {
	
	r := mux.NewRouter();
	r = web.Up(*r);

	http.ListenAndServe(":3000", r)





	
	


	

}
