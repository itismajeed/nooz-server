package app


//App is app container
type App struct{
	Server *Server
	Config *Config
}

// New app
func  New() *App  {

	config := &Config{
		Name: "Nooz",
		ResponseTimeOut:15,
		RequestTimeOut:15,
		AppAddr:"localhost:8000",
	}
	
	server :=&Server{
		Addr:config.AppAddr,
		Router:getRoutes(),
	}

	app :=&App{
		Config:config,
		Server:server,
	}
	return app
}