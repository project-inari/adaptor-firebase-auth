package service

import (
	"context"
	"errors"
	"testing"

	"github.com/project-inari/adaptor-firebase-auth/dto"
	"github.com/stretchr/testify/assert"
)

type mockFirebaseAuthRepository struct {
	token string
	err   error
}

func (m *mockFirebaseAuthRepository) SignUp(_ context.Context, _ dto.SignUpReq, _ dto.SignUpReqHeader) (string, error) {
	return m.token, m.err
}

const (
	mockToken    = "mockToken"
	mockUsername = "mockUsername"
	mockEmail    = "mock@email.com"
	mockPassword = "mockPassword"
	mockPhoneNo  = "+66000000000"

	enAcceptLocale = "EN"
	thAcceptLocale = "TH"
)

func TestSignUp(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {

		svc := New(Dependencies{
			FirebaseAuthRepository: &mockFirebaseAuthRepository{
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
		assert.Equal(t, mockToken, res.Token)
	})

	t.Run("error", func(t *testing.T) {
		svc := New(Dependencies{
			FirebaseAuthRepository: &mockFirebaseAuthRepository{
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
		assert.Equal(t, "", res.Token)
	})
}
