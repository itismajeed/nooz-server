package main

import (
	"nooz-server/app"
)

func main(){
	app:=app.New()
	app.Server.ServeHTTP()
}