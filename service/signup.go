package service

import (
	"context"

	"github.com/project-inari/adaptor-firebase-auth/dto"
)

// SignUp creates a new user in Firebase Auth
func (s *service) SignUp(ctx context.Context, req dto.SignUpReq, header dto.SignUpReqHeader) (*dto.SignUpRes, error) {
	res, err := s.firebaseAuthRepository.SignUp(ctx, req, header)
	if err != nil {
		return nil, err
	}

	return &dto.SignUpRes{
		Username: req.Username,
		UID:      res.UID,
		Token:    res.Token,
	}, nil
}
