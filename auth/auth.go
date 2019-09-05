package auth

import (
	"context"
	firebase "firebase.google.com/go"
	"log"
	"net/http"
)

func WithAuth(handlerFunc http.HandlerFunc) http.Handler {
	return Authenticated(handlerFunc)
}

func Authenticated(next http.Handler) http.Handler {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		client, err := app.Auth(ctx)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}

		token, err := client.VerifyIDToken(ctx, r.Header.Get("Authorization"))
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r.Header.Set("X-User-UID", token.UID)

		log.Printf("Verified ID token: %v\n", token)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
