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

// SignUpInfo represents the information of the signed up user
type SignUpInfo struct {
	UID   string
	Token string
}

// VerifyTokenInfo represents the information of the verified token
type VerifyTokenInfo struct {
	Username string
	UID      string
}

// SignUp creates a new user in Firebase Auth
func (r *firebaseAuthRepository) SignUp(ctx context.Context, payload dto.SignUpReq, header dto.SignUpReqHeader) (*SignUpInfo, error) {
	params := (&auth.UserToCreate{}).
		Email(payload.Email).
		Password(payload.Password).
		DisplayName(payload.Username).
		PhoneNumber(payload.PhoneNo)

	user, err := r.client.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error - [firebaseAuthRepository.SignUp] unable to create user: %v", err)
	}

	err = r.client.SetCustomUserClaims(ctx, user.UID, map[string]interface{}{
		"username":      payload.Username,
		"accept-locale": header.AcceptLocale,
	})
	if err != nil {
		return nil, fmt.Errorf("error - [firebaseAuthRepository.SignUp] unable to set user claims: %v", err)
	}

	user, err = r.client.GetUser(ctx, user.UID)
	if err != nil {
		return nil, fmt.Errorf("error - [firebaseAuthRepository.SignUp] unable to get user: %v", err)
	}

	token, err := r.client.CustomTokenWithClaims(ctx, user.UID, map[string]interface{}{
		"username":      payload.Username,
		"accept-locale": header.AcceptLocale,
	})
	if err != nil {
		return nil, fmt.Errorf("error - [firebaseAuthRepository.SignUp] unable to generate token: %v", err)
	}

	return &SignUpInfo{
		UID:   user.UID,
		Token: token,
	}, nil
}

// VerifyToken verifies the token with Firebase Auth
func (r *firebaseAuthRepository) VerifyToken(ctx context.Context, token string) (*VerifyTokenInfo, error) {
	t, err := r.client.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("error - [firebaseAuthRepository.VerifyToken] unable to verify token: %v", err)
	}

	username, ok := t.Claims["username"].(string)
	if !ok {
		return nil, fmt.Errorf("error - [firebaseAuthRepository.VerifyToken] username claim is not a string")
	}

	return &VerifyTokenInfo{
		Username: username,
		UID:      t.UID,
	}, nil
}

// DeleteUser deletes a user in Firebase Auth
func (r *firebaseAuthRepository) DeleteUser(ctx context.Context, uid string) error {
	err := r.client.DeleteUser(ctx, uid)
	if err != nil {
		return fmt.Errorf("error - [firebaseAuthRepository.DeleteUser] unable to delete user: %v", err)
	}

	return nil
}
