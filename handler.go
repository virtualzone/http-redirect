package main

import (
	"net/http"
	"strings"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	target := GetConfig().TargetURL
	if GetConfig().AppendPath {
		path := strings.TrimPrefix(r.URL.Path, "/")
		if path != "" {
			target += "/" + path
		}
	}
	w.Header().Set("Location", target)
	w.WriteHeader(GetConfig().HttpStatusCode)
}
