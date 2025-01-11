package service

import (
	"context"
	"log/slog"

	"github.com/project-inari/adaptor-firebase-auth/dto"
)

// SignUp creates a new user in Firebase Auth
func (s *service) SignUp(ctx context.Context, req dto.SignUpReq, header dto.SignUpReqHeader) (*dto.SignUpRes, error) {
	res, err := s.firebaseAuthRepository.SignUp(ctx, req, header)
	if err != nil {
		slog.Error("error - [SignUp] unable to sign up", slog.Any("error", err))
		return &dto.SignUpRes{}, err
	}

	slog.Info("[firebaseAuthRepository.SignUp] user created successfully", slog.Any("username", req.Username))

	return &dto.SignUpRes{
		Token: res,
	}, nil
}
