package di

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"

	"github.com/project-inari/adaptor-firebase-auth/config"
)

func setupFirebaseAuth(pctx context.Context, c config.FirebaseAuthConfig) (*auth.Client, error) {
	firebaseConfig := &firebase.Config{
		ProjectID:        c.ProjectID,
		ServiceAccountID: c.ServiceAccountID,
	}

	fb, err := firebase.NewApp(pctx, firebaseConfig, option.WithCredentialsJSON([]byte(c.CredentialsJSON)))
	if err != nil {
		return nil, err
	}

	client, err := fb.Auth(pctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
