package oauth

import (
	"context"
	"fmt"
	"github.com/hellodoctordev/common/firebase"
	"github.com/hellodoctordev/common/keys"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"log"
)

var (
	firestoreClient         = firebase.NewFirestoreClient()
	googleOAuthDoctorConfig = &oauth2.Config{
		ClientID:     keys.GoogleOAuthKeys.DoctorClientID,
		ClientSecret: keys.GoogleOAuthKeys.DoctorClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.google.com/o/oauth2/auth",
			TokenURL:  "https://oauth2.googleapis.com/token",
			AuthStyle: 0,
		},
	}
	googleOAuthPatientConfig = &oauth2.Config{
		ClientID:     keys.GoogleOAuthKeys.PatientClientID,
		ClientSecret: keys.GoogleOAuthKeys.PatientClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.google.com/o/oauth2/auth",
			TokenURL:  "https://oauth2.googleapis.com/token",
			AuthStyle: 0,
		},
	}
)

func GetGoogleUserOAuthTokenSource(userID string, userOAuthCode ...string) (tokenSource option.ClientOption, err error) {
	type StoredOAuthToken struct {
		Token    oauth2.Token
		Provider string
	}

	ctx := context.Background()

	userSnapshot, err := firestoreClient.Doc(fmt.Sprintf("users/%s", userID)).Get(ctx)
	if err != nil {
		log.Printf("error occurred getting user snapshot: %s", err)
		return
	}

	role, _ := userSnapshot.DataAt("account.role")

	var oauthConfig  *oauth2.Config

	if role.(string) == "provider" {
		oauthConfig = googleOAuthDoctorConfig
	} else {
		oauthConfig = googleOAuthPatientConfig
	}

	userOAuthCollectionRef := firestoreClient.Collection(fmt.Sprintf("users/%s/oauthTokens", userID))

	var oauthToken StoredOAuthToken

	userOAuth, err := userOAuthCollectionRef.
		Where("Provider", "==", "google").
		Documents(ctx).
		Next()

	if err != nil && len(userOAuthCode) == 1 {
		token, err2 := oauthConfig.Exchange(ctx, userOAuthCode[0], oauth2.AccessTypeOffline)
		if err2 != nil {
			log.Printf("error exchanging oauth token: %s", err2)
			return
		}

		oauthToken := StoredOAuthToken{
			Token:    *token,
			Provider: "google",
		}

		_, _, err2 = userOAuthCollectionRef.Add(ctx, oauthToken)
		if err2 != nil {
			log.Printf("error storing user oauth: %s", err2)
			return
		}

	} else if err != nil {
		log.Printf("error getting user oauth: %s", err)
		return

	} else {
		err2 := userOAuth.DataTo(&oauthToken)
		if err2 != nil {
			log.Printf("error parsing user oauth: %s", err)
			return
		}
	}

	return option.WithTokenSource(oauthConfig.TokenSource(ctx, &oauthToken.Token)), nil
}
