package config

import (
	"log"
	"net/http"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

var (
	App *Application
	Mux *http.ServeMux
)

func Init(infoLog, errorLog *log.Logger) {

	App = &Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	Mux = http.NewServeMux()
}
