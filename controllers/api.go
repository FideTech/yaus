package controllers

import (
	"net/http"
	"strings"

	"github.com/FideTech/yaus/core"
	"github.com/FideTech/yaus/models"
	"github.com/FideTech/yaus/storage"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Replace(r.URL.Path, "/api/v1", "", 1)

	if strings.HasPrefix(path, "/short-links") {
		shortLinksHandler(w, r, path)
		return
	}
}

func shortLinksHandler(w http.ResponseWriter, r *http.Request, path string) {
	subPath := strings.Replace(path, "/short-links", "", 1)

	if subPath != "" {
		shortLinkHandler(w, r, subPath)
		return
	}

	switch r.Method {
	case http.MethodGet:
		links, err := core.GetAllDynamicShortLinks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sendJSON(w, links)
		return
	case http.MethodPost:
		var shortLink models.ShortLink
		if err := bindJSON(r, &shortLink); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := core.CreateDynamicShortLink(&shortLink); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sendJSON(w, shortLink)
		return
	}
}

func shortLinkHandler(w http.ResponseWriter, r *http.Request, subPath string) {
	key := strings.Replace(subPath, "/", "", 1)

	switch r.Method {
	case http.MethodGet:
		link, err := core.GetDynamicShortLinkByKey(key)
		if err != nil {
			if err == storage.ErrNotFound {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sendJSON(w, link)
		return
	default:
		http.Error(w, "not implemented", http.StatusNotImplemented)
		return
	}
}
