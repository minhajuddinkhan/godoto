package main

import (
	"net/http"

	"github.com/minhajuddinkhan/godoto/web"
)

func main() {

	web.Up()

	http.ListenAndServe(":3000", nil)
}
