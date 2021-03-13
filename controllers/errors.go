package controllers

import (
	"net/http"

	"github.com/FideTech/yaus/config"
)

//ErrorHandler handles the hard coded "/e/" short links
func ErrorHandler(w http.ResponseWriter, r *http.Request, value string) {
	link, found := config.Config.Hardcoded.Errors[value]
	if !found {
		http.Error(w, "link not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, link.URL, http.StatusTemporaryRedirect)
}
