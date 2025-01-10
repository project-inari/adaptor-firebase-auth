package repository

import "firebase.google.com/go/auth"

type firebaseAuthRepository struct {
	client *auth.Client
}

type FirebaseAuthRepositoryDependencies struct {
	Client *auth.Client
}

// NewFirebaseAuthRepository creates a new instance of firebaseAuthRepository
func NewFirebaseAuthRepository(d FirebaseAuthRepositoryDependencies) *firebaseAuthRepository {
	return &firebaseAuthRepository{client: d.Client}
}
