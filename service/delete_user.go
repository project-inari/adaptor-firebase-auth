package service

import (
	"context"

	"github.com/project-inari/adaptor-firebase-auth/dto"
)

func (s *service) DeleteUser(ctx context.Context, req dto.DeleteUserReq) (*dto.DeleteUserRes, error) {
	if err := s.firebaseAuthRepository.DeleteUser(ctx, req.UID); err != nil {
		return nil, err
	}

	return &dto.DeleteUserRes{
		Success: true,
	}, nil
}
