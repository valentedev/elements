package main

import (
	"log"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	mb := mailbox{
		"status": "available",
		"system_info": map[string]string{
			"enviroment": app.config.env,
			"version":    "1.0",
		},
	}

	err := app.writeJSON(w, http.StatusOK, mb, nil)
	if err != nil {
		log.Println(err)
	}
}
