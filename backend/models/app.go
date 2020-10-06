package models

import "github.com/gorilla/mux"

// Application instance contains all runtime configurations and internal structures
type Application struct {
	Config ApplicationConfiguration
	Router *mux.Router
}

// ApplicationConfiguration contains all the ENV variables that will be used within the backend
type ApplicationConfiguration struct {
	Port string
}
