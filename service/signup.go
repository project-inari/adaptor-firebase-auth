package service

import (
	"context"

	"github.com/project-inari/adaptor-firebase-auth/dto"
)

// SignUp creates a new user in Firebase Auth
func (s *service) SignUp(ctx context.Context, req dto.SignUpReq, header dto.SignUpReqHeader) (*dto.SignUpRes, error) {
	res, err := s.firebaseAuthRepository.SignUp(ctx, req, header)
	if err != nil {
		return &dto.SignUpRes{}, err
	}

	return &dto.SignUpRes{
		Token: res,
	}, nil
}
