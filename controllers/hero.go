package controllers

import (
	"net/http"
	"fmt"
)

func Health(w http.ResponseWriter, r *http.Request) {
	
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "health200");
}
