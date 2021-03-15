package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func sendJSON(w http.ResponseWriter, data interface{}) {
	marshalled, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}

func bindJSON(r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}
