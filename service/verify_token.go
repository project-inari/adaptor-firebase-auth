package service

import (
	"context"
	"log/slog"

	"github.com/project-inari/adaptor-firebase-auth/dto"
)

func (s *service) VerifyToken(ctx context.Context, req dto.VerifyTokenReq) *dto.VerifyTokenRes {
	res, err := s.firebaseAuthRepository.VerifyToken(ctx, req.Token)
	if err != nil {
		slog.Error("error - [VerifyToken] unable to verify token", slog.Any("error", err))
		return &dto.VerifyTokenRes{Success: false}
	}

	slog.Info("[firebaseAuthRepository.VerifyToken] token verified successfully", slog.Any("username", res.Username))

	return &dto.VerifyTokenRes{
		Username: res.Username,
		UID:      res.UID,
		Success:  true,
	}
}
