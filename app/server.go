package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)
// Server is a struct to serve http requests
type Server struct{
	Addr string
	Router *mux.Router
}


// ServeHTTP serves http
func (s *Server) ServeHTTP()  {
	fmt.Printf("Server started on %v...\n",s.Addr)
	http.ListenAndServe(s.Addr,s.Router)
	
} 	