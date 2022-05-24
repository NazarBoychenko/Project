package http

import (
	"log"
	"net/http"
)

func NotFoundBD() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal("You have problem with link to BD!")
	}
}
