// Package service provides the business logic service layer for the server
package service

import (
	"github.com/project-inari/adaptor-firebase-auth/repository"
)

// Port represents the service layer functions
type Port interface {
}

type service struct {
	firebaseAuthRepository repository.FirebaseAuthRepository
}

// Dependencies represents the dependencies for the service
type Dependencies struct {
	FirebaseAuthRepository repository.FirebaseAuthRepository
}

// New creates a new service
func New(d Dependencies) Port {
	return &service{
		firebaseAuthRepository: d.FirebaseAuthRepository,
	}
}
