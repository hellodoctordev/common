package auth

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/hellodoctordev/common/keys"
	"log"
	"net/http"
	"strings"
)

func WithAuth(handlerFunc http.HandlerFunc) http.Handler {
	return Authenticated(handlerFunc)
}

func WithAuthRoles(handlerFunc http.HandlerFunc, roles... string) http.Handler {
	return Authenticated(WithRole(handlerFunc, roles...))
}

func WithInternalAuth(handlerFunc http.HandlerFunc) http.Handler {
	return AuthenticatedInternalService(handlerFunc)
}

func WithAdminAuth(handlerFunc http.HandlerFunc) http.Handler {
	return AuthenticatedAdmin(handlerFunc)
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

		token, err := client.VerifyIDToken(ctx, getRequestAuthorizationToken(r))
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r.Header.Set("X-User-ID", token.UID)

		if token.Claims["role"] != nil {
			r.Header.Set("X-User-Role", token.Claims["role"].(string))
		}

		if scopesArray, ok := token.Claims["scopes"].([]string); ok {
			r.Header.Set("X-User-Scopes", strings.Join(scopesArray, ","))
		}

		if token.Claims["groupID"] != nil {
			r.Header.Set("X-User-Group-ID", token.Claims["groupID"].(string))
		}

		if token.Claims["authorizedRemoteUserID"] != nil {
			remoteUserUID := token.Claims["authorizedRemoteUserID"].(string)
			
			r.Header.Set("X-Remote-User-ID", remoteUserUID)
		}

		next.ServeHTTP(w, r)
	})
}

func WithRole(next http.Handler, roles... string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestRole := r.Header.Get("X-User-Role")

		for _, role := range roles {
			if requestRole == role {
				next.ServeHTTP(w, r)
				return
			}
		}

		log.Printf("role '%s' not authorized", requestRole)
		w.WriteHeader(http.StatusUnauthorized)
		return
	})
}

func AuthenticatedInternalService(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("X-Internal-Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(keys.InternalServiceKeys.ServiceSecret), nil
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

func AuthenticatedAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("X-Internal-Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(keys.AdminKeys.AdminTokenSecret), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if ok := token.Claims.(jwt.MapClaims).VerifyIssuer(keys.AdminKeys.AdminTokenIssuer, true); !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getRequestAuthorizationToken(r *http.Request) string {
	authorizationHeader := r.Header.Get("Authorization")

	if strings.Contains(authorizationHeader, "Bearer") {
		return strings.Split(authorizationHeader, " ")[1]
	} else {
		return authorizationHeader
	}
}
