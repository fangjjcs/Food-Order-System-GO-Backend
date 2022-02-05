package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {

	wrapper := make(map[string]interface{})

	wrapper["status"] = status
	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status int) {

	type jsonError struct {
		Message string `json:"message"`
	}
	theError := jsonError{
		Message: err.Error(),
	}
	app.writeJSON(w, status, theError, "error")
}
