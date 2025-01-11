// Package repository provides the repository interfaces for the domain
package repository

import (
	"context"

	"github.com/project-inari/adaptor-firebase-auth/dto"
)

// FirebaseAuthRepository represents the repository functions of Firebase Auth
type FirebaseAuthRepository interface {
	SignUp(ctx context.Context, payload dto.SignUpReq, header dto.SignUpReqHeader) (string, error)
	VerifyToken(ctx context.Context, token string) (*VerifyTokenInfo, error)
}
