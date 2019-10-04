package handlers

import (
	"net/http"
)
//HomeHandler handles / root requests
func HomeHandler(w http.ResponseWriter, r *http.Request)  {
	welcomeText := "Hello welcome to noozApp!"
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type","application/json")
	w.Write([]byte(welcomeText))
}