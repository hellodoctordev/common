package auth

import (
	"github.com/hellodoctordev/common/utils"
	"net/http"
	"strings"
)

func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Accept")

		origin := r.Header.Get("Origin")

		allowedOrigins := []string{
			"https://hellodoctor-staging-cast.firebaseapp.com",
			"http://api.stage.hellodoctor.com.mx",
			"http://api.hellodoctor.com.mx",
			"https://cast.hellodoctor.com.mx",
		}

		if utils.ContainsString(allowedOrigins, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// FIXME Only allow localhost/ngrok.io in DEV deployment
		if origin == "http://localhost:3000" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else if strings.HasSuffix(origin, "ngrok.io") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
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
