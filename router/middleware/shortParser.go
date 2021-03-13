package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

type ShortLinkHandler func(http.ResponseWriter, *http.Request, string)

type ParseShortLink struct {
	handler ShortLinkHandler
	prefix  string
}

func (psl *ParseShortLink) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	short := strings.Replace(r.URL.Path, fmt.Sprintf("/%s/", psl.prefix), "", 1)

	if short == "" {
		http.Error(w, "link not found", http.StatusNotFound)
		return
	}

	psl.handler(w, r, short)
}

func NewParsedLink(handlerToWrap ShortLinkHandler, prefix string) *ParseShortLink {
	return &ParseShortLink{handlerToWrap, prefix}
}
