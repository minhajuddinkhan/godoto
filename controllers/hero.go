package controllers

import (
	"net/http"
	"io"
)

func Hero(w http.ResponseWriter, r *http.Request) {
	

	resp, err := http.Get("https://api.opendota.com/api/heroes")
	if err != nil {
        // handle error
        return
    }
    defer resp.Body.Close()
    io.Copy(w, resp.Body)
}
