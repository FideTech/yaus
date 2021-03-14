package controllers

import (
	"net/http"

	"github.com/FideTech/yaus/config"
)

//InfoHandler handles the hard coded "/i/" short links
func InfoHandler(w http.ResponseWriter, r *http.Request, value string) {
	link, found := config.Config.Hardcoded.Infos[value]
	if !found {
		http.Error(w, "link not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, link.URL, http.StatusTemporaryRedirect)
}
