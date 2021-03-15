package controllers

import (
	"log"
	"net/http"

	"github.com/FideTech/yaus/core"
)

//DynamicLinkHandler handles the dynamically created short links at the endpoint "/d/"
func DynamicLinkHandler(w http.ResponseWriter, r *http.Request, value string) {
	link, err := core.GetDynamicShortLinkByKey(value)
	if err != nil {
		http.Error(w, "link not found", http.StatusNotFound)
		return
	}

	if err := core.AddRedirectToDynamicShortLink(value); err != nil {
		log.Printf("failed to add redirect count for %s", value)
	}

	http.Redirect(w, r, link.URL, http.StatusTemporaryRedirect)
}
