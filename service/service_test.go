package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/project-inari/adaptor-firebase-auth/dto"
	"github.com/project-inari/adaptor-firebase-auth/repository"
)

type mockFirebaseAuthRepository struct {
	token string
	err   error

	username string
	uid      string
}

func (m *mockFirebaseAuthRepository) SignUp(_ context.Context, _ dto.SignUpReq, _ dto.SignUpReqHeader) (*repository.SignUpInfo, error) {
	return &repository.SignUpInfo{
		UID:   mockUID,
		Token: m.token,
	}, m.err
}

func (m *mockFirebaseAuthRepository) VerifyToken(_ context.Context, _ string) (*repository.VerifyTokenInfo, error) {
	return &repository.VerifyTokenInfo{
		Username: m.username,
		UID:      m.uid,
	}, m.err
}

func (m *mockFirebaseAuthRepository) DeleteUser(_ context.Context, _ string) error {
	return m.err
}

const (
	mockToken    = "mockToken"
	mockUsername = "mockUsername"
	mockEmail    = "mock@email.com"
	mockPassword = "mockPassword"
	mockPhoneNo  = "+66000000000"
	mockUID      = "mockUID"

	enAcceptLocale = "EN"
	thAcceptLocale = "TH"
)

func TestSignUp(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {

		svc := New(Dependencies{
			FirebaseAuthRepository: &mockFirebaseAuthRepository{
				uid:   mockUID,
				token: mockToken,
				err:   nil,
			},
		})

		req := dto.SignUpReq{
			Username: mockUsername,
			Email:    mockEmail,
			Password: mockPassword,
			PhoneNo:  mockPhoneNo,
		}
		header := dto.SignUpReqHeader{
			AcceptLocale: enAcceptLocale,
		}

		res, err := svc.SignUp(ctx, req, header)

		assert.Nil(t, err)
		assert.Equal(t, mockUsername, res.Username)
		assert.Equal(t, mockUID, res.UID)
		assert.Equal(t, mockToken, res.Token)
	})

	t.Run("error", func(t *testing.T) {
		svc := New(Dependencies{
			FirebaseAuthRepository: &mockFirebaseAuthRepository{
				uid:   "",
				token: "",
				err:   errors.New("error"),
			},
		})

		req := dto.SignUpReq{
			Username: mockUsername,
			Email:    mockEmail,
			Password: mockPassword,
			PhoneNo:  mockPhoneNo,
		}

		header := dto.SignUpReqHeader{
			AcceptLocale: thAcceptLocale,
		}

		res, err := svc.SignUp(ctx, req, header)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}

func TestVerifyToken(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		svc := New(Dependencies{
			FirebaseAuthRepository: &mockFirebaseAuthRepository{
				username: mockUsername,
				uid:      mockUID,
				err:      nil,
			},
		})

		req := dto.VerifyTokenReq{
			Token: mockToken,
		}

		res := svc.VerifyToken(ctx, req)

		assert.Equal(t, mockUsername, res.Username)
		assert.Equal(t, mockUID, res.UID)
		assert.True(t, res.Success)
	})

	t.Run("error", func(t *testing.T) {
		svc := New(Dependencies{
			FirebaseAuthRepository: &mockFirebaseAuthRepository{
				username: "",
				uid:      "",
				err:      errors.New("error"),
			},
		})

		req := dto.VerifyTokenReq{
			Token: mockToken,
		}

		res := svc.VerifyToken(ctx, req)

		assert.Equal(t, "", res.Username)
		assert.Equal(t, "", res.UID)
		assert.False(t, res.Success)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		svc := New(Dependencies{
			FirebaseAuthRepository: &mockFirebaseAuthRepository{
				err: nil,
			},
		})

		req := dto.DeleteUserReq{
			UID: mockUID,
		}

		res, err := svc.DeleteUser(ctx, req)

		assert.Nil(t, err)
		assert.True(t, res.Success)
	})

	t.Run("error", func(t *testing.T) {
		ctx := context.Background()

		svc := New(Dependencies{
			FirebaseAuthRepository: &mockFirebaseAuthRepository{
				err: errors.New("error"),
			},
		})

		req := dto.DeleteUserReq{
			UID: mockUID,
		}

		res, err := svc.DeleteUser(ctx, req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}
