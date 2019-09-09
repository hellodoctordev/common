package auth

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/hellodoctordev/common/keys"
	"log"
	"net/http"
)

func WithAuth(handlerFunc http.HandlerFunc) http.Handler {
	return Authenticated(handlerFunc)
}

func WithAuthRole(role string, handlerFunc http.HandlerFunc) http.Handler {
	return Authenticated(WithRole(role, handlerFunc))
}

func WithInternalAuth(handlerFunc http.HandlerFunc) http.Handler {
	return AuthenticatedInternalService(handlerFunc)
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
		r.Header.Set("X-User-Role", token.Claims["role"].(string))

		next.ServeHTTP(w, r)
	})
}

func WithRole(role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestRole := r.Header.Get("X-User-Role")

		if requestRole != role {
			log.Printf("role '%s' not authorized", requestRole)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthenticatedInternalService(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("X-Internal-Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return keys.InternalServiceKeys.ServiceSecret, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
