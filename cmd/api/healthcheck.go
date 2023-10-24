package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	response, err := json.Marshal(data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and couldn't process your request", http.StatusInternalServerError)
		return
	}

	response = append(response, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.Write(response)
}
