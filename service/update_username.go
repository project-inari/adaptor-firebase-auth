package service

import (
	"context"

	"github.com/project-inari/adaptor-firebase-auth/dto"
)

// UpdateUsername updates the username of the user in Firebase Auth
func (s *service) UpdateUsername(ctx context.Context, req dto.UpdateUsernameReq) (*dto.UpdateUsernameRes, error) {
	err := s.firebaseAuthRepository.UpdateUsername(ctx, req.UID, req.NewUsername)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateUsernameRes{
		Username: req.NewUsername,
		UID:      req.UID,
		Success:  true,
	}, nil
}
