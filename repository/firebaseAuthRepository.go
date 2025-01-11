package repository

import (
	"context"
	"fmt"

	"firebase.google.com/go/auth"

	"github.com/project-inari/adaptor-firebase-auth/dto"
)

type firebaseAuthRepository struct {
	client *auth.Client
}

// FirebaseAuthRepositoryDependencies represents the dependencies for the firebaseAuthRepository
type FirebaseAuthRepositoryDependencies struct {
	Client *auth.Client
}

// NewFirebaseAuthRepository creates a new instance of firebaseAuthRepository
func NewFirebaseAuthRepository(d FirebaseAuthRepositoryDependencies) FirebaseAuthRepository {
	return &firebaseAuthRepository{client: d.Client}
}

// SignUp creates a new user in Firebase Auth
func (r *firebaseAuthRepository) SignUp(ctx context.Context, payload dto.SignUpReq, header dto.SignUpReqHeader) (string, error) {
	params := (&auth.UserToCreate{}).
		Email(payload.Email).
		Password(payload.Password).
		DisplayName(payload.Username).
		PhoneNumber(payload.PhoneNo)

	user, err := r.client.CreateUser(ctx, params)
	if err != nil {
		return "", fmt.Errorf("error - [firebaseAuthRepository.SignUp] unable to create user: %v", err)
	}

	err = r.client.SetCustomUserClaims(ctx, user.UID, map[string]interface{}{
		"username":      payload.Username,
		"accept-locale": header.AcceptLocale,
	})
	if err != nil {
		return "", fmt.Errorf("error - [firebaseAuthRepository.SignUp] unable to set user claims: %v", err)
	}

	user, err = r.client.GetUser(ctx, user.UID)
	if err != nil {
		return "", fmt.Errorf("error - [firebaseAuthRepository.SignUp] unable to get user: %v", err)
	}

	token, err := r.client.CustomTokenWithClaims(ctx, user.UID, map[string]interface{}{
		"username":      payload.Username,
		"accept-locale": header.AcceptLocale,
	})
	if err != nil {
		return "", fmt.Errorf("error - [firebaseAuthRepository.SignUp] unable to generate token: %v", err)
	}

	return token, nil
}
