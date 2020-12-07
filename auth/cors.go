package auth

import (
	"net/http"
	"strings"
)

func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Accept")

		origin := r.Header.Get("Origin")
		referer := r.Header.Get("Referer")

		// FIXME Only allow localhost/ngrok.io in DEV deployment
		if origin == "http://localhost:3000" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else if strings.HasSuffix(origin, "ngrok.io") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else if origin == "http://api.stage.hellodoctor.com.mx" {
			w.Header().Set("Access-Control-Allow-Origin", "api.stage.hellodoctor.com.mx")
		} else if origin == "http://api.hellodoctor.com.mx" {
			w.Header().Set("Access-Control-Allow-Origin", "api.hellodoctor.com.mx")
		} else if origin == "https://cast.hellodoctor.com.mx" || referer == "https://cast.hellodoctor.com.mx" {
			w.Header().Set("Access-Control-Allow-Origin", "cast.hellodoctor.com.mx")
		}

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func WithCORSFunc(next http.HandlerFunc) http.Handler {
	return WithCORS(next)
}
