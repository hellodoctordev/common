package auth

import "net/http"

func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Accept")

		// FIXME Only allow localhost in DEV deployment
		if r.Header.Get("Origin") == "http://localhost:3000" {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		} else if r.Header.Get("Origin") == "http://api.stage.hellodoctor.com.mx" {
			w.Header().Set("Access-Control-Allow-Origin", "api.stage.hellodoctor.com.mx")
		}

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
