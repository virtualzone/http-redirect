package main

import (
	"log"
	"net/http"
	"strings"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	target := GetConfig().TargetURL
	if GetConfig().AppendPath {
		if path != "/" {
			target += path
		}
	}
	log.Printf("%d redirect: %s -> %s\n", GetConfig().HttpStatusCode, path, target)
	w.Header().Set("Location", target)
	w.WriteHeader(GetConfig().HttpStatusCode)
}
