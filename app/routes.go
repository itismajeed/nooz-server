package app

import(
	"github.com/gorilla/mux"
	"nooz-server/app/handlers"
	
)
func getRoutes() *mux.Router{
	r:= mux.NewRouter();
	r.HandleFunc("/",handlers.HomeHandler)
	return r
}