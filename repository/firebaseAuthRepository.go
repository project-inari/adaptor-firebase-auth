package repository

import "firebase.google.com/go/auth"

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
