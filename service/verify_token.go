package service

import (
	"context"

	"github.com/project-inari/adaptor-firebase-auth/dto"
)

func (s *service) VerifyToken(ctx context.Context, req dto.VerifyTokenReq) *dto.VerifyTokenRes {
	res, err := s.firebaseAuthRepository.VerifyToken(ctx, req.Token)
	if err != nil {
		return &dto.VerifyTokenRes{Success: false}
	}

	return &dto.VerifyTokenRes{
		Username: res.Username,
		UID:      res.UID,
		Success:  true,
	}
}
